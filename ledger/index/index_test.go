/*
 *    Copyright 2018 INS Ecosystem
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package index

import (
	"os"
	"testing"

	"github.com/insolar/insolar/ledger/record"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestDecodeLifelineReturnsCorrectResult(t *testing.T) {
	idx := Lifeline{
		LatestStateID: record.ID{1},
		AppendIDs:     []record.ID{{1}, {2}},
	}

	encoded := EncodeLifeline(&idx)
	assert.Equal(
		t,
		[]byte{0xa2, 0x69, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x73, 0x82, 0x58, 0x20, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x58, 0x20, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6d, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x58, 0x20, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		encoded,
	)

	decoded := DecodeLifeline(encoded)
	assert.Equal(t, decoded, idx)
}
