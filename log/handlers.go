//
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
//

package log

import (
	"fmt"
	"net/http"

	"github.com/insolar/insolar/insolar"
	"github.com/rs/zerolog"
)

type loglevelChangeHandler struct {
}

func NewLoglevelChangeHandler() http.Handler {
	handler := &loglevelChangeHandler{}
	return handler
}

// ServeHTTP is an HTTP handler that changes the global minimum log level
func (h *loglevelChangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	levelStr := "(nil)"
	if values["level"] != nil {
		levelStr = values["level"][0]
	}
	level, err := insolar.ParseLevel(levelStr)
	if err != nil {
		w.WriteHeader(500)
		_, _ = fmt.Fprintf(w, "Invalid level '%v': %v\n", levelStr, err)
		return
	}

	zlevel, err := InternalLevelToZerologLevel(level)
	if err != nil {
		w.WriteHeader(500)
		_, _ = fmt.Fprintf(w, "Invalid level '%v': %v\n", levelStr, err)
		return
	}

	zerolog.SetGlobalLevel(zlevel)

	w.WriteHeader(200)
	_, _ = fmt.Fprintf(w, "New log level: '%v'\n", levelStr)
}
