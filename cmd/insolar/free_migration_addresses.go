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

package main

import (
	"fmt"

	"github.com/insolar/insolar/api/sdk"
	"github.com/pkg/errors"
)

var shardsAtOneTime = 10

func getfreeMigrationCount(adminUrls []string, publicUrls []string, memberKeysDirPath string, shardsCount int, alertCount int) {
	insSDK, err := sdk.NewSDK(adminUrls, publicUrls, memberKeysDirPath)
	check("SDK is not initialized: ", err)

	var shoudAlert []map[int]int
	var freeAdressesInShards = map[int]int{}
	for i := 0; i < shardsCount; i += shardsAtOneTime {
		part, err := insSDK.GetAddressesCount(i)
		check(fmt.Sprintf("Error while getting addresses from shards %d to %d: ", i, i+shardsAtOneTime), err)
		partSliced, ok := part.([]interface{})
		if !ok {
			check(fmt.Sprintf("Error while getting addresses from shards %d to %d: ", i, i+shardsAtOneTime), errors.New("error while converting result"))
		}
		// var migrationShardsMap = map[int]int{}
		for _, r := range partSliced {
			rMap := r.(map[string]interface{})
			s, ok := rMap["shardIndex"].(float64)
			if !ok {
				check(
					fmt.Sprintf("Error while getting addresses from shards %d to %d: ", i, i+shardsAtOneTime),
					errors.New("error while converting shardIndex"),
				)
			}
			shardIndex := int(s)

			f, ok := rMap["freeAmount"].(float64)
			if !ok {
				check(
					fmt.Sprintf("Error while getting addresses from shard %d: ", shardIndex),
					errors.New("error while converting freeAmount"),
				)
			}
			freeAmount := int(f)

			freeAdressesInShards[shardIndex] = freeAmount
			if freeAmount <= alertCount {
				shoudAlert = append(shoudAlert, map[int]int{shardIndex: freeAmount})
			}
		}
	}
	fmt.Println("free addresses by shard: ", freeAdressesInShards)
	if len(shoudAlert) > 0 {
		fmt.Println("ALERT: too little free addresses in shards: ", shoudAlert)
	}
}
