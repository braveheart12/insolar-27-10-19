///
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
///

package testutils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// nodesErrorLogReader reads launchnet logs and trying to find errors it them
// when errors found notification writes to 'writer'.
// It reads files in different goroutines and return chan for stopping them.
func NodesErrorLogReader(logsDir string, writer io.Writer) chan struct{} {
	closeChan := make(chan struct{})
	wg := sync.WaitGroup{}

	logs, err := getLogs(logsDir)
	if err != nil {
		writeToOutput(writer, fmt.Sprintf("Can't find node logs: %s \n", err))
	}

	wg.Add(len(logs))
	for _, fileName := range logs {
		fName := fileName // be careful using loops and values in parallel code
		go readLogs(writer, &wg, fName, closeChan)
	}

	wg.Wait()
	return closeChan
}

func getLogs(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == "output.log" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func readLogs(writer io.Writer, wg *sync.WaitGroup, fileName string, closeChan chan struct{}) {
	defer wg.Done()

	file, err := os.Open(fileName)
	if err != nil {
		writeToOutput(writer, fmt.Sprintln("Can't open log file ", fileName, ", error : ", err))
	}
	_, err = file.Seek(-1, io.SeekEnd)
	if err != nil {
		writeToOutput(writer, fmt.Sprintln("Can't seek through log file ", fileName, ", error : ", err))
	}

	// for making wg.Done()
	go findErrorsInLog(writer, fileName, file, closeChan)
}

func findErrorsInLog(writer io.Writer, fName string, file io.ReadCloser, closeChan chan struct{}) {
	defer file.Close()
	reader := bufio.NewReader(file)

	ok := true
	for ok {
		select {
		case <-time.After(time.Millisecond):
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				writeToOutput(writer, fmt.Sprintln("Can't read string from ", fName, ", error: ", err))
				ok = false
			}

			if strings.Contains(line, " ERR ") {
				writeToOutput(writer, fmt.Sprintln("!!! THERE ARE ERRORS IN ERROR LOG !!! ", fName))
				ok = false
			}
		case <-closeChan:
			ok = false
		}

	}
}

// TODO rewrite code from goroutines to save log files size before test
// TODO and read it after test in same goroutine
// TODO in current realization if check() cause exit we get leaked goroutine
func writeToOutput(out io.Writer, data string) {
	_, err := out.Write([]byte(data))
	check("Can't write data to output", err)
}

func check(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
