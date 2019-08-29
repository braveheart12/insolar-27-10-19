///
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
///

package main

import (
	"github.com/insolar/insolar/conveyor/smachine"
	"github.com/insolar/insolar/conveyor/smachine/main/example"
	"time"
)

func main() {
	sm := smachine.NewSlotMachine(smachine.SlotMachineConfig{
		SlotPageSize:  10,
		PollingPeriod: 100 * time.Millisecond,
	})

	sm.AddNew(smachine.NoLink(), &example.StateMachine1{})

	for i := 0; ; i++ {
		sm.ScanOnce(nil)
		//fmt.Printf("%03d %v: slots=%v\n", i, time.Now(), sm.OccupiedSlotCount())
		time.Sleep(10 * time.Millisecond)
	}
}