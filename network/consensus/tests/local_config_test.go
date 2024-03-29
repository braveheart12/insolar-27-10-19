//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package tests

import (
	"context"
	"time"

	"github.com/insolar/insolar/network/consensus/common/cryptkit"
	"github.com/insolar/insolar/network/consensus/gcpv2/api"
)

const defaultNshGenerationDelay = time.Millisecond * 0

var RoundTimingsFor1s = api.RoundTimings{
	StartPhase0At: 100 * time.Millisecond, // Not scaled

	StartPhase1RetryAt: 00 * time.Millisecond, // 0 for no retries
	EndOfPhase1:        250 * time.Millisecond,
	EndOfPhase2:        400 * time.Millisecond,
	EndOfPhase3:        500 * time.Millisecond,
	EndOfConsensus:     600 * time.Millisecond,

	BeforeInPhase2ChasingDelay: 0 * time.Millisecond,
	BeforeInPhase3ChasingDelay: 0 * time.Millisecond,
}

var EphemeralRoundTimingsFor1s = api.RoundTimings{
	StartPhase0At: 100 * time.Millisecond, // Not scaled

	StartPhase1RetryAt: 0 * time.Millisecond, // 0 for no retries
	EndOfPhase1:        200 * time.Millisecond,
	EndOfPhase2:        600 * time.Millisecond,
	EndOfPhase3:        800 * time.Millisecond,
	EndOfConsensus:     900 * time.Millisecond,

	BeforeInPhase2ChasingDelay: 0 * time.Millisecond,
	BeforeInPhase3ChasingDelay: 0 * time.Millisecond,
}

func NewEmuLocalConfig(ctx context.Context) api.LocalNodeConfiguration {
	r := emuLocalConfig{timings: RoundTimingsFor1s, ephemeralTimings: EphemeralRoundTimingsFor1s, ctx: ctx}
	return &r
}

type emuLocalConfig struct {
	timings          api.RoundTimings
	ephemeralTimings api.RoundTimings
	ctx              context.Context
}

func (r *emuLocalConfig) GetNodeCountHint() int {
	return 10
}

func (r *emuLocalConfig) GetParentContext() context.Context {
	return r.ctx
}

func (r *emuLocalConfig) PublicKeyStore() {
}

func (r *emuLocalConfig) AsPublicKeyStore() cryptkit.PublicKeyStore {
	return r
}

func (r *emuLocalConfig) PrivateKeyStore() {
}

func (r *emuLocalConfig) getConsensusTimings(t api.RoundTimings, nextPulseDelta uint16) api.RoundTimings {
	if nextPulseDelta == 1 {
		return t
	}
	m := time.Duration(nextPulseDelta) // this is NOT a duration, but a multiplier

	t.StartPhase0At *= 1 // don't scale!
	t.StartPhase1RetryAt *= m
	t.EndOfPhase1 *= m
	t.EndOfPhase2 *= m
	t.EndOfPhase3 *= m
	t.EndOfConsensus *= m
	t.BeforeInPhase2ChasingDelay *= m
	t.BeforeInPhase3ChasingDelay *= m

	return t
}

func (r *emuLocalConfig) GetConsensusTimings(nextPulseDelta uint16) api.RoundTimings {
	return r.getConsensusTimings(r.timings, nextPulseDelta)
}

func (r *emuLocalConfig) GetEphemeralTimings(nextPulseDelta uint16) api.RoundTimings {
	return r.getConsensusTimings(r.ephemeralTimings, nextPulseDelta)
}

func (r *emuLocalConfig) GetSecretKeyStore() cryptkit.SecretKeyStore {
	return r
}
