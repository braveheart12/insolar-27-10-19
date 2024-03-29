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

package goplugintestutils

import (
	"go/build"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

const insolarImportPath = "github.com/insolar/insolar"

var testdataDir = testdataPath()

func buildCLI(name string) (string, error) {
	binPath := filepath.Join(testdataDir, name)

	out, err := exec.Command(
		"go", "build",
		"-o", binPath,
		filepath.Join(insolarImportPath, name),
	).CombinedOutput()
	if err != nil {
		return "", errors.Wrapf(err, "can't build preprocessor. buildPrototypes output: %s", string(out))
	}
	return binPath, nil
}

func buildInsiderCLI() (string, error) {
	return buildCLI("cmd/insgorund")
}

func BuildPreprocessor() (string, error) {
	return buildCLI("application/cmd/insgocc")
}

func testdataPath() string {
	p, err := build.Default.Import("github.com/insolar/insolar", "", build.FindOnly)
	if err != nil {
		panic(err)
	}
	return filepath.Join(p.Dir, "testdata", "logicrunner")
}

// Build compiles and return path to insgorund and insgocc binaries.
func Build() (string, string, error) {
	icc, err := buildInsiderCLI()
	if err != nil {
		return "", "", err
	}

	insgocc, err := BuildPreprocessor()
	if err != nil {
		return "", "", err
	}
	return icc, insgocc, nil
}
