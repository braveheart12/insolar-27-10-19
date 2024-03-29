package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/dgraph-io/badger"

	"github.com/gorilla/mux"
	"github.com/insolar/insolar/log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type MergeJSONRequest struct {
	BkpName string `json:"bkpName"`
	RunGC   bool   `json:"runGC"`
}

type MergeJSONResponse struct {
	Message string `json:"message"`
}

var globalBadgerHandler *badger.DB
var globalBadgerLock sync.Mutex // see comments below

func sendHTTPResponse(w http.ResponseWriter, statusCode int, resp MergeJSONResponse) {
	h := w.Header()
	h.Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("sendHttpResonse, json.Marshal: %v", err)
		return
	}

	log.Infof("sendHttpResonse: statusCode = %d, resp = %s", statusCode, respBytes)

	_, err = w.Write(respBytes)
	if err != nil {
		log.Errorf("sendHttpResonse, w.Write: %v", err)
	}
}

func MergeHTTPHandler(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("MergeHTTPHandler, ioutil.ReadAll: %v", err)
		return
	}

	log.Infof("Processing request: %s", reqBytes)

	var req MergeJSONRequest
	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Errorf("MergeHTTPHandler, json.Unmarshal: %v", err)
		return
	}

	if req.BkpName == "" {
		sendHTTPResponse(w, 400, MergeJSONResponse{
			Message: "Missing bkpName",
		})
		return
	}

	log.Infof("Merging incremental backup, bkpName = %s", req.BkpName)

	msg := "Merge done."
	ok := func() bool { // Note: defer works on function level, not scope level!
		bkpFile, err := os.Open(req.BkpName)
		if err != nil {
			sendHTTPResponse(w, 400, MergeJSONResponse{
				Message: "Failed to open incremental backup file",
			})
			return false
		}
		defer bkpFile.Close()

		log.Info("Backup file is opened")
		if globalBadgerHandler == nil {
			sendHTTPResponse(w, 500, MergeJSONResponse{
				Message: "DB handler is nil",
			})
			return false
		}

		// only one globalBadgerHandler.Load() and GC can run at a time!
		log.Info("MergeHTTPHandler - about to Lock() globalBadgerLock")
		globalBadgerLock.Lock()
		log.Info("MergeHTTPHandler - globalBadgerLock Locked(), executing  globalBadgerHandler.Load()")
		defer func() {
			log.Info("MergeHTTPHandler - calling globalBadgerLock.Unlock()")
			globalBadgerLock.Unlock()
		}()

		err = globalBadgerHandler.Load(bkpFile, 1)
		if err != nil {
			sendHTTPResponse(w, 400, MergeJSONResponse{
				Message: "Failed to load incremental backup file",
			})
			return false
		}

		if req.RunGC {
			log.Info("Running GC...")
			err = globalBadgerHandler.RunValueLogGC(0.7)
			var m string
			if err == nil {
				m = " GC done."
			} else {
				m = " GC failed: " + err.Error()
			}
			log.Info(m)
			msg += m
		}

		return true
	}()

	if !ok {
		return
	}

	sendHTTPResponse(w, 200, MergeJSONResponse{
		Message: msg,
	})
}

func daemon(listenAddr string, targetDBPath string) {
	log.Info("Opening DB...")

	ops := badger.DefaultOptions(targetDBPath)
	ops.Logger = badgerLogger
	var err error
	func() { // Note: defer works on function level, not scope level!
		log.Info("daemon() - about to Lock() globalBadgerLock")
		// it guarantees defined behavior if terms of `globalBadgerHandler` value
		globalBadgerLock.Lock()
		log.Info("daemon() - globalBadgerLock Locked(), executing  globalBadgerHandler.Load()")
		defer func() {
			log.Info("daemon() - calling globalBadgerLock.Unlock()")
			globalBadgerLock.Unlock()
		}()

		// Note: make sure `globalBadgerHandler` will be assigned (don't use := here)
		globalBadgerHandler, err = badger.Open(ops)
		if err != nil {
			err = errors.Wrap(err, "failed to open DB")
			exitWithError(err)
		}
		log.Info("DB opened")

		err = isDBEmpty(globalBadgerHandler)
		if err == nil {
			// will exit
			err = errors.New("DB must not be empty")
			closeRawDB(globalBadgerHandler, err)
		}
	}()

	log.Info("DB opened and it's not empty. Starting HTTP server...")
	r := mux.NewRouter().
		PathPrefix("/api/v1").
		Path("/merge").
		Subrouter()
	r.Methods("POST").
		HandlerFunc(MergeHTTPHandler)
	http.Handle("/", r)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
	log.Info("HTTP server terminated\n")
}

func parseDaemonParams() *cobra.Command {
	var (
		listenAddr   string
		targetDBPath string
	)

	var daemonCmd = &cobra.Command{
		Use:   "daemon",
		Short: "run merge daemon",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("Starting merge daemon, address = %s, target-db = %s", listenAddr, targetDBPath)
			daemon(listenAddr, targetDBPath)
		},
	}
	mergeFlags := daemonCmd.Flags()
	targetDBFlagName := "target-db"
	mergeFlags.StringVarP(
		&targetDBPath, targetDBFlagName, "t", "", "directory where backup will be roll to (required)")
	mergeFlags.StringVarP(
		&listenAddr, "address", "a", ":8080", "listen address")

	err := cobra.MarkFlagRequired(mergeFlags, targetDBFlagName)
	if err != nil {
		err := errors.Wrap(err, "failed to set required param: "+targetDBFlagName)
		exitWithError(err)
	}

	return daemonCmd
}

func daemonMerge(address string, backupFileName string, runGC bool) {
	reqJSON := MergeJSONRequest{BkpName: backupFileName, RunGC: runGC}
	reqBytes, err := json.Marshal(reqJSON)
	if err != nil {
		err = errors.Wrap(err, "daemonMerge - json.Marshal failed")
		exitWithError(err)
	}

	req, err := http.NewRequest("POST", address+"/api/v1/merge", bytes.NewBuffer(reqBytes))
	if err != nil {
		err = errors.Wrap(err, "daemonMerge - http.NewRequest failed")
		exitWithError(err)
	}

	client := http.Client{}
	httpResp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "daemonMerge - client.Do failed")
		exitWithError(err)
	}
	defer httpResp.Body.Close()

	respBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		err = errors.Wrap(err, "daemonMerge - ioutil.ReadAll failed")
		exitWithError(err)
	}
	var resp MergeJSONResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		err = errors.Wrap(err, "daemonMerge - json.Unmarshal failed")
		exitWithError(err)
	}

	if httpResp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Merge failed: daemon returned code %d and body: %s", httpResp.StatusCode, respBytes))
		exitWithError(err)
	}

	log.Infof("HTTP response OK. Daemon: %s", resp.Message)
}

func parseDaemonMergeParams() *cobra.Command {
	var (
		address        string
		backupFileName string
		runGC          bool
	)

	var daemonMergeCmd = &cobra.Command{
		Use:   "daemon-merge",
		Short: "merge incremental backup using merge daemon",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("Starting daemon-merge, address = %s, bkp-name = %s", address, backupFileName)
			daemonMerge(address, backupFileName, runGC)
		},
	}
	mergeFlags := daemonMergeCmd.Flags()
	bkpFileName := "bkp-name"
	mergeFlags.StringVarP(
		&backupFileName, bkpFileName, "n", "", "file name if incremental backup (required)")
	mergeFlags.StringVarP(
		&address, "address", "a", "http://localhost:8080", "merge daemon listen address")
	mergeFlags.BoolVarP(
		&runGC, "run-gc", "g", false, "run GC")

	err := cobra.MarkFlagRequired(mergeFlags, bkpFileName)
	if err != nil {
		err := errors.Wrap(err, "failed to set required param: "+bkpFileName)
		exitWithError(err)
	}

	return daemonMergeCmd
}
