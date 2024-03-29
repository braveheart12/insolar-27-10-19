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

package ph01ctl

import (
	"context"
	"fmt"
	"time"

	"github.com/insolar/insolar/network/consensus/gcpv2/api/member"
	"github.com/insolar/insolar/network/consensus/gcpv2/core/coreapi"
	"github.com/insolar/insolar/network/consensus/gcpv2/core/population"
	"github.com/insolar/insolar/network/consensus/gcpv2/phasebundle/pulsectl"

	"github.com/insolar/insolar/network/consensus/common/endpoints"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/phases"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/transport"
	"github.com/insolar/insolar/network/consensus/gcpv2/core"
)

func NewPhase01PrepController(s pulsectl.PulseSelectionStrategy) *Phase01PrepController {
	return &Phase01PrepController{pulseStrategy: s}
}

var _ core.PrepPhaseController = &Phase01PrepController{}

type Phase01PrepController struct {
	core.PrepPhaseControllerTemplate
	core.HostPacketDispatcherTemplate

	realm         *core.PrepRealm
	pulseStrategy pulsectl.PulseSelectionStrategy
}

func (c *Phase01PrepController) CreatePacketDispatcher(pt phases.PacketType, realm *core.PrepRealm) population.PacketDispatcher {
	c.realm = realm
	return c
}

func (c *Phase01PrepController) GetPacketType() []phases.PacketType {
	return []phases.PacketType{phases.PacketPhase0, phases.PacketPhase1}
}

func (c *Phase01PrepController) DispatchHostPacket(ctx context.Context, packet transport.PacketParser,
	from endpoints.Inbound, flags coreapi.PacketVerifyFlags) error {

	var pp transport.PulsePacketReader
	var nr member.Rank

	switch packet.GetPacketType() {
	case phases.PacketPhase0:
		p0 := packet.GetMemberPacket().AsPhase0Packet()
		nr = p0.GetNodeRank()
		pp = p0.GetEmbeddedPulsePacket()
	case phases.PacketPhase1:
		p1 := packet.GetMemberPacket().AsPhase1Packet()
		// if p1.HasFullIntro() || p1.HasCloudIntro() || p1.HasJoinerSecret() {
		//	return fmt.Errorf("introduction data were not expected: from=%v", from)
		// }
		nr = p1.GetAnnouncementReader().GetNodeRank()
		if p1.HasPulseData() {
			pp = p1.GetEmbeddedPulsePacket()
		}
	default:
		panic("illegal value")
	}
	if nr.IsJoiner() && pp != nil {
		return fmt.Errorf("pulse data in Phase0/Phas1 is not allowed from a joiner: from=%v", from)
	}
	// if { TODO check ranks? }

	ok, err := c.pulseStrategy.HandlePulsarPacket(ctx, pp, from, false)
	if err != nil || !ok {
		return err
	}

	startedAt := time.Now()
	return c.realm.ApplyPulseData(ctx, startedAt, pp, false, from)
}
