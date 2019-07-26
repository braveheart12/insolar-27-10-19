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

package inspectors

import (
	"context"
	"strings"

	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/profiles"
	"github.com/insolar/insolar/network/consensus/gcpv2/phasebundle/nodeset"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/network/consensus/common/cryptkit"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/member"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/proofs"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/statevector"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/transport"
	"github.com/insolar/insolar/network/consensus/gcpv2/core"
)

type VectorInspectionFactory interface {
	CreateVectorInspection(inlineLimit int) VectorInspection
}

type VectorInspection interface {
	CreateInspector(scanner nodeset.VectorEntryScanner, digestFactory transport.ConsensusDigestFactory,
		nodeID insolar.ShortNodeID) VectorInspector
}

type VectorInspector interface {
	GetBitset() member.StateBitset
	/* Must be called before any CreateVector or InspectVector, and before any parallel access */
	PrepareForInspection(ctx context.Context)
	CreateVector(cryptkit.DigestSigner) statevector.Vector
	InspectVector(ctx context.Context, sender *core.NodeAppearance, customOptions uint32, otherData statevector.Vector) InspectedVector
	CreateNextPopulation(nodeset.ConsensusBitsetRow) ([]profiles.PopulationRank, proofs.CloudStateHash, proofs.GlobulaStateHash)
}

type InspectedVector interface {
	IsInspected() bool
	HasMissingMembers() bool
	HasSenderFault() bool
	Reinspect(ctx context.Context, inspector VectorInspector) InspectedVector

	Inspect(ctx context.Context)
	GetInspectionResults() (*nodeset.ConsensusStatRow, nodeset.NodeVerificationResult)

	GetBitset() member.StateBitset
	GetNode() *core.NodeAppearance
	GetCustomOptions() uint32
}

func NewVectorInspectionFactory() VectorInspectionFactory {
	return &vectorInspectionFactory{}
}

type vectorInspectionFactory struct {
}

func (v vectorInspectionFactory) CreateVectorInspection(inlineLimit int) VectorInspection {
	return NewVectorInspection(inlineLimit)
}

func NewVectorInspection(maxPopulationForInlineHashing int) VectorInspection {
	return &vectorInspection{maxPopulationForInlineHashing}
}

func NewIgnorantVectorInspection() VectorInspection {
	return &ignorantVectorInspection{}
}

type vectorInspection struct {
	maxPopulationForInlineHashing int
}

func (p vectorInspection) CreateInspector(scanner nodeset.VectorEntryScanner, digestFactory transport.ConsensusDigestFactory,
	nodeID insolar.ShortNodeID) VectorInspector {

	r := &vectorInspectorImpl{VectorBuilder: nodeset.NewVectorBuilder(digestFactory, scanner,
		make(member.StateBitset, scanner.GetIndexedCount())),

		maxPopulationForInlineHashing: p.maxPopulationForInlineHashing,
		nodeID:                        nodeID}

	r.FillBitset()
	return r
}

type ignorantVectorInspection struct {
}

func (p ignorantVectorInspection) CreateInspector(scanner nodeset.VectorEntryScanner, digestFactory transport.ConsensusDigestFactory,
	nodeID insolar.ShortNodeID) VectorInspector {

	r := &vectorIgnorantInspectorImpl{VectorBuilder: nodeset.NewVectorBuilder(digestFactory, scanner,
		make(member.StateBitset, scanner.GetIndexedCount())),
		nodeID: nodeID}

	r.FillBitset()
	return r
}

type vectorIgnorantInspectorImpl struct {
	nodeID insolar.ShortNodeID
	nodeset.VectorBuilder
}

func (p *vectorIgnorantInspectorImpl) CreateNextPopulation(selectionSet nodeset.ConsensusBitsetRow) ([]profiles.PopulationRank, proofs.CloudStateHash, proofs.GlobulaStateHash) {

	return createNextPopulation(&p.VectorBuilder, selectionSet)
}

func (p *vectorIgnorantInspectorImpl) PrepareForInspection(ctx context.Context) {
}

func (p *vectorIgnorantInspectorImpl) CreateVector(signer cryptkit.DigestSigner) statevector.Vector {
	panic("illegal state")
}

func (p *vectorIgnorantInspectorImpl) InspectVector(ctx context.Context, sender *core.NodeAppearance, customOptions uint32,
	otherVector statevector.Vector) InspectedVector {

	if p.GetBitset().Len() != otherVector.Bitset.Len() {
		panic("illegal state - StateBitset length mismatch")
	}

	r := ignoredVector{parent: p, node: sender, otherData: otherVector, customOptions: customOptions}
	r.comparedStats = nodeset.CompareToStatRow(p.GetBitset(), otherVector.Bitset)
	r.verifyResult = nodeset.NvrNotVerified

	if otherVector.Trusted.AnnouncementHash == nil && otherVector.Doubted.AnnouncementHash == nil {
		r.verifyResult |= nodeset.NvrSenderFault
	}

	if r.comparedStats.HasValues(nodeset.ComparedMissingHere) {
		// we can't validate anything without data
		// ...  check for updates or/and send requests
		r.verifyResult |= nodeset.NvrMissingNodes
	}

	vr := nodeset.NvrTrustedValid | nodeset.NvrDoubtedValid
	sr := nodeset.SummarizeStats(r.otherData.Bitset, vr, r.comparedStats)
	sr.SetCustomOptions(r.customOptions)
	r.nodeStats = &sr

	return &r
}

type ignoredVector struct {
	parent        *vectorIgnorantInspectorImpl
	node          *core.NodeAppearance
	customOptions uint32
	otherData     statevector.Vector
	verifyResult  nodeset.NodeVerificationResult
	comparedStats nodeset.ComparedBitsetRow
	nodeStats     *nodeset.ConsensusStatRow
}

func (p *ignoredVector) GetCustomOptions() uint32 {
	return p.customOptions
}

func (p *ignoredVector) GetNode() *core.NodeAppearance {
	return p.node
}

func (p *ignoredVector) GetInspectionResults() (*nodeset.ConsensusStatRow, nodeset.NodeVerificationResult) {
	return p.nodeStats, p.verifyResult
}

func (p *ignoredVector) GetBitset() member.StateBitset {
	return p.otherData.Bitset
}

func (p *ignoredVector) String() string {
	switch p.verifyResult {
	case nodeset.NvrNotVerified, nodeset.NvrSenderFault:
		return p.verifyResult.String()
	}

	b := strings.Builder{}
	b.WriteByte('[')
	p.verifyResult.StringPart(&b)
	b.WriteString("]∑")
	b.WriteString(p.comparedStats.StringSummary())

	return b.String()
}

func (p *ignoredVector) IsInspected() bool {
	return true
}

func (p *ignoredVector) HasMissingMembers() bool {
	return p.verifyResult&nodeset.NvrMissingNodes != 0
}

func (p *ignoredVector) HasSenderFault() bool {
	return p.verifyResult&nodeset.NvrSenderFault != 0
}

func (p *ignoredVector) Reinspect(ctx context.Context, inspector VectorInspector) InspectedVector {
	if newParent, ok := inspector.(*vectorIgnorantInspectorImpl); ok && p.parent == newParent {
		return p
	}
	return inspector.InspectVector(ctx, p.node, p.customOptions, p.otherData)
}

func (p *ignoredVector) Inspect(ctx context.Context) {
	// do nothing, all was done before
}

type vectorInspectorImpl struct {
	nodeID                        insolar.ShortNodeID
	maxPopulationForInlineHashing int
	Trusted                       statevector.CalcSubVector
	Doubted                       statevector.CalcSubVector
	nodeset.VectorBuilder
}

func (p *vectorInspectorImpl) CreateNextPopulation(selectionSet nodeset.ConsensusBitsetRow) ([]profiles.PopulationRank, proofs.CloudStateHash, proofs.GlobulaStateHash) {

	return createNextPopulation(&p.VectorBuilder, selectionSet)
}

func (p *vectorInspectorImpl) ensureHashes() {
	if p.Trusted.AnnouncementHash == nil {
		panic("illegal state")
	}
}

func (p *vectorInspectorImpl) PrepareForInspection(ctx context.Context) {
	if p.Trusted.AnnouncementHash != nil {
		panic("illegal state")
	}

	p.Trusted.AnnouncementHash, p.Doubted.AnnouncementHash = p.BuildAllGlobulaAnnouncementHashes()

	p.Trusted.CalcStateWithRank, p.Doubted.CalcStateWithRank =
		p.BuildGlobulaStateHashesAndRanks(true, p.Doubted.AnnouncementHash != nil, p.nodeID,
			statevector.CalcStateWithRank{}, statevector.CalcStateWithRank{})
}

func (p *vectorInspectorImpl) CreateVector(signer cryptkit.DigestSigner) statevector.Vector {
	p.ensureHashes()
	return statevector.NewVector(p.GetBitset(), p.Trusted.Sign(signer), p.Doubted.Sign(signer))
}

func (p *vectorInspectorImpl) InspectVector(ctx context.Context, sender *core.NodeAppearance, customOptions uint32,
	otherVector statevector.Vector) InspectedVector {

	if p.GetBitset().Len() != otherVector.Bitset.Len() {
		panic("illegal state - StateBitset length mismatch")
	}
	p.ensureHashes()

	r := newInspectedVectorAndPreInspect(ctx, p, sender, customOptions, otherVector)
	return &r
}

func newInspectedVectorAndPreInspect(ctx context.Context, p *vectorInspectorImpl, sender *core.NodeAppearance,
	customOptions uint32, otherVector statevector.Vector) inspectedVector {

	r := inspectedVector{parent: p, node: sender, otherData: otherVector, customOptions: customOptions}
	r.comparedStats = nodeset.CompareToStatRow(p.GetBitset(), otherVector.Bitset)

	r.verifyResult = nodeset.NvrNotVerified

	if otherVector.Trusted.AnnouncementHash == nil && otherVector.Doubted.AnnouncementHash == nil {
		r.verifyResult |= nodeset.NvrSenderFault
	}

	if r.comparedStats.HasValues(nodeset.ComparedMissingHere) {
		// we can't validate anything without data
		// ...  check for updates or/and send requests
		r.verifyResult |= nodeset.NvrMissingNodes
		return r
	}

	r.trustedPart, r.doubtedPart = nodeset.PrepareSubVectorsComparison(r.comparedStats,
		otherVector.Trusted.AnnouncementHash != nil,
		otherVector.Doubted.AnnouncementHash != nil)

	if p.maxPopulationForInlineHashing == 0 || p.maxPopulationForInlineHashing >= p.GetEntryScanner().GetIndexedCount() {
		r.Inspect(ctx)
	}
	return r
}

type inspectedVector struct {
	parent                   *vectorInspectorImpl
	node                     *core.NodeAppearance
	otherData                statevector.Vector
	customOptions            uint32
	trustedPart, doubtedPart nodeset.SubVectorCompared
	verifyResult             nodeset.NodeVerificationResult
	comparedStats            nodeset.ComparedBitsetRow
	nodeStats                *nodeset.ConsensusStatRow
}

func (p *inspectedVector) GetCustomOptions() uint32 {
	return p.customOptions
}

func (p *inspectedVector) GetNode() *core.NodeAppearance {
	return p.node
}

func (p *inspectedVector) GetInspectionResults() (*nodeset.ConsensusStatRow, nodeset.NodeVerificationResult) {
	return p.nodeStats, p.verifyResult
}

func (p *inspectedVector) GetBitset() member.StateBitset {
	return p.otherData.Bitset
}

func (p *inspectedVector) String() string {
	switch p.verifyResult {
	case nodeset.NvrNotVerified, nodeset.NvrSenderFault:
		return p.verifyResult.String()
	}

	b := strings.Builder{}
	b.WriteByte('[')
	p.verifyResult.StringPart(&b)
	b.WriteString("]∑")
	b.WriteString(p.comparedStats.StringSummary())

	return b.String()
}

func (p *inspectedVector) verifySignature(localHash proofs.GlobulaStateHash, remoteSignature cryptkit.SignatureHolder) bool {
	return localHash != nil && p.node.GetSignatureVerifier().IsValidDigestSignature(localHash, remoteSignature)
}

func (p *inspectedVector) IsInspected() bool {
	return p.verifyResult != nodeset.NvrNotVerified
}

func (p *inspectedVector) HasMissingMembers() bool {
	return p.verifyResult&nodeset.NvrMissingNodes != 0
}

func (p *inspectedVector) HasSenderFault() bool {
	return p.verifyResult&nodeset.NvrSenderFault != 0
}

func (p *inspectedVector) Reinspect(ctx context.Context, inspector VectorInspector) InspectedVector {
	if newParent, ok := inspector.(*vectorInspectorImpl); ok && p.parent == newParent {
		return p
	}
	return inspector.InspectVector(ctx, p.node, p.customOptions, p.otherData)
}

func (p *inspectedVector) Inspect(ctx context.Context) {

	nodeID := p.node.GetNodeID()

	if nodeID == 0 || p.IsInspected() {
		return
	}
	if p.nodeStats != nil {
		panic("illegal state")
	}

	p.verifyResult = p.doVerifyVectorHashes(ctx)
	if p.verifyResult == nodeset.NvrNotVerified {
		panic("illegal state")
	}

	vr := p.verifyResult &^ nodeset.NvrHashlessFlags
	if p.verifyResult == nodeset.NvrNotVerified {
		panic("illegal state")
	}

	ns := nodeset.SummarizeStats(p.otherData.Bitset, vr, p.comparedStats)
	ns.SetCustomOptions(p.customOptions)

	p.nodeStats = &ns
}

func (p *inspectedVector) doVerifyVectorHashes(ctx context.Context) nodeset.NodeVerificationResult {

	selfData := statevector.CalcVector{Trusted: p.parent.Trusted, Doubted: p.parent.Doubted}

	if p.doubtedPart.IsNeeded() && selfData.Doubted.AnnouncementHash == nil {
		// special case when all our nodes are in trusted, so other's doubted vector will be matched with the trusted one of ours
		selfData.Doubted = selfData.Trusted
	}

	gahTrusted, gahDoubted := selfData.Trusted.AnnouncementHash, selfData.Doubted.AnnouncementHash

	var vectorBuilder nodeset.VectorBuilder //
	if p.trustedPart.IsRecalc() || p.doubtedPart.IsRecalc() {
		vectorBuilder = p.parent.CreateDerived(p.comparedStats)

		// It does remap the original bitset with the given stats
		gahTrusted, gahDoubted = vectorBuilder.BuildGlobulaAnnouncementHashes(
			p.trustedPart.IsRecalc(), p.doubtedPart.IsRecalc(), gahTrusted, gahDoubted)
	}

	log := inslogger.FromContext(ctx)

	validTrusted := p.trustedPart.IsNeeded() && gahTrusted.Equals(p.otherData.Trusted.AnnouncementHash)
	validDoubted := p.doubtedPart.IsNeeded() && gahDoubted.Equals(p.otherData.Doubted.AnnouncementHash)

	if log.Is(insolar.DebugLevel) {
		if validTrusted != p.trustedPart.IsNeeded() || validDoubted != p.doubtedPart.IsNeeded() {
			log.Errorf("mismatched AnnouncementHash:\n Here: %v %v\nThere: %v %v",
				gahTrusted, gahDoubted, p.otherData.Trusted.AnnouncementHash, p.otherData.Doubted.AnnouncementHash)
		}
	}

	verifyRes := nodeset.NvrNotVerified
	if validDoubted && !validTrusted {
		// As Trusted is a subset of Doubted, then Doubted can't be valid if Trusted is not.
		// This is an evident fraud/error by the sender.
		// Use status for doubted, but ignore results for Trusted check
		// TODO report fraud
		verifyRes |= nodeset.NvrSenderFault
		p.trustedPart = nodeset.SvcIgnore
	}

	// TODO make a controlled parameter!

	//if validTrusted || validDoubted {
	//	recalcTrusted := p.trustedPart.IsRecalc() && validTrusted
	//	recalcDoubted := p.doubtedPart.IsRecalc() && validDoubted
	//
	//	gshTrusted, gshDoubted := selfData.Trusted.CalcStateWithRank, selfData.Doubted.CalcStateWithRank
	//	if recalcTrusted || recalcDoubted {
	//		gshTrusted, gshDoubted = vectorBuilder.BuildGlobulaStateHashesAndRanks(recalcTrusted, recalcDoubted,
	//			p.node.GetNodeID(), gshTrusted, gshDoubted)
	//
	//		if recalcTrusted {
	//			validTrusted = gshTrusted.ExpectedRank == p.otherData.Trusted.ExpectedRank
	//		}
	//		if recalcDoubted {
	//			validDoubted = gshDoubted.ExpectedRank == p.otherData.Doubted.ExpectedRank
	//		}
	//
	//		if log.Is(insolar.DebugLevel) {
	//			if recalcTrusted && !validTrusted || recalcDoubted && !validDoubted {
	//				log.Errorf("mismatched ExpectedRank:\n Here: %v %v\nThere: %v %v",
	//					gshTrusted.ExpectedRank, gshDoubted.ExpectedRank,
	//					p.otherData.Trusted.ExpectedRank, p.otherData.Doubted.ExpectedRank)
	//			}
	//		}
	//	}
	//
	//	prevValidTrusted := validTrusted
	//	prevValidDoubted := validDoubted
	//
	//	validTrusted = validTrusted && p.verifySignature(gshTrusted.StateHash, p.otherData.Trusted.StateSignature)
	//	validDoubted = validDoubted && p.verifySignature(gshDoubted.StateHash, p.otherData.Doubted.StateSignature)
	//
	//	if log.Is(insolar.DebugLevel) {
	//		if validTrusted != prevValidTrusted || validDoubted != prevValidDoubted {
	//			log.Errorf("mismatched signature of StateHash:\n Here: %v %v\nThere: %v %v",
	//				gshTrusted.StateHash, gshDoubted.StateHash,
	//				p.otherData.Trusted.StateSignature, p.otherData.Doubted.StateSignature)
	//		}
	//	}
	//}

	if p.trustedPart.IsNeeded() {
		verifyRes.SetTrusted(validTrusted, p.trustedPart.IsRecalc())
	}
	if p.doubtedPart.IsNeeded() {
		verifyRes.SetDoubted(validDoubted, p.doubtedPart.IsRecalc())
	}

	return verifyRes
}

func createNextPopulation(p *nodeset.VectorBuilder, selectionSet nodeset.ConsensusBitsetRow) ([]profiles.PopulationRank, proofs.CloudStateHash, proofs.GlobulaStateHash) {

	result := make([]profiles.PopulationRank, p.GetEntryScanner().GetSortedCount())
	newIndex := 0
	gshRank := p.BuildGlobulaStateHashWithFilter(insolar.AbsentShortNodeID,
		func(nodeData nodeset.VectorEntryData, postponed bool, filter uint32) {

			mode := member.OpMode(filter)
			if mode.IsPowerless() && !postponed {
				panic("illegal state")
			}

			result[newIndex].Profile = nodeData.Profile
			result[newIndex].OpMode = mode
			if mode.IsPowerless() {
				result[newIndex].Power = 0
			} else {
				result[newIndex].Power = nodeData.RequestedPower
			}

			newIndex++
		},
		func(index int, nodeData nodeset.VectorEntryData, parentFilter uint32) (bool, uint32) {

			newMode := nodeData.RequestedMode
			if index >= 0 { // not a joiner
				decision := selectionSet.Get(index)
				if nodeData.IsEmpty() { // we can't really cope with it ... so lets pretend we can
					decision = nodeset.CbsSuspected
				}
				if decision != nodeset.CbsIncluded {
					newMode = consensusDecisionToOpMode(decision, nodeData.Profile.GetOpMode())
				}
			}
			return nodeData.RequestedPower == 0 || newMode.IsPowerless(), uint32(newMode)
		})

	return result[:newIndex], gshRank.StateHash, gshRank.StateHash
}

func consensusDecisionToOpMode(decision nodeset.ConsensusBitsetEntry, lastMode member.OpMode) member.OpMode {
	switch decision {
	case nodeset.CbsFraud:
		if lastMode.IsMistrustful() {
			return member.ModeEvictedAsFraud
		}
		if lastMode.IsSuspended() {
			return member.ModePossibleFraudAndSuspected
		}
		return member.ModePossibleFraud
	case nodeset.CbsSuspected:
		if lastMode.IsSuspended() {
			return member.ModeEvictedAsSuspected
		}
		if lastMode.IsMistrustful() {
			return member.ModePossibleFraudAndSuspected
		}
		return member.ModeSuspected
	case nodeset.CbsExcluded:
		return member.ModeEvictedAsFraud
	default:
		panic("illegal value")
	}
}
