/*
 * The Clear BSD License
 *
 * Copyright (c) 2019 Insolar Technologies
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted (subject to the limitations in the disclaimer below) provided that
 * the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *  * Neither the name of Insolar Technologies nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
 * BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND
 * CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING,
 * BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package merkle

import (
	"sort"

	"github.com/insolar/insolar/core"
	"github.com/pkg/errors"
)

type PulseEntry struct {
	Pulse *core.Pulse
}

func (pe *PulseEntry) hash(helper *merkleHelper) []byte {
	return helper.pulseHash(pe.Pulse)
}

type GlobuleEntry struct {
	*PulseEntry
	ProofSet      map[core.Node]*PulseProof
	PulseHash     []byte
	PrevCloudHash []byte
	GlobuleID     core.GlobuleID
}

func (ge *GlobuleEntry) hash(helper *merkleHelper) ([]byte, error) {
	nodeEntryByRole := nodeEntryByRole(ge.PulseEntry, ge.ProofSet)
	var bucketHashes [][]byte

	for _, role := range core.AllStaticRoles {
		roleEntries, ok := nodeEntryByRole[role]
		if !ok {
			continue
		}

		sortEntries(roleEntries)
		bucketEntryRoot, err := roleEntryRoot(roleEntries, helper)

		if err != nil {
			return nil, errors.Wrap(err, "[ hash ] Failed to create tree for bucket role entry")
		}

		bucketInfoHash := helper.bucketInfoHash(role, uint32(len(roleEntries)))
		bucketHash := helper.bucketHash(bucketInfoHash, bucketEntryRoot)
		bucketHashes = append(bucketHashes, bucketHash)
	}

	tree, err := treeFromHashList(bucketHashes, helper.scheme.IntegrityHasher())

	if err != nil {
		return nil, errors.Wrap(err, "[ hash ] Failed to create tree for bucket hashes")
	}

	return tree.Root(), nil
}

type CloudEntry struct {
	ProofSet      []*GlobuleProof
	PrevCloudHash []byte
}

func (ce *CloudEntry) hash(helper *merkleHelper) ([]byte, error) {
	var result [][]byte

	for _, proof := range ce.ProofSet {
		globuleInfoHash := helper.globuleInfoHash(ce.PrevCloudHash, uint32(proof.GlobuleID), proof.NodeCount)
		globuleHash := helper.globuleHash(globuleInfoHash, proof.NodeRoot)
		result = append(result, globuleHash)
	}

	tree, err := treeFromHashList(result, helper.scheme.IntegrityHasher())
	if err != nil {
		return nil, errors.Wrap(err, "[ hash ] Failed to create tree")
	}

	return tree.Root(), nil
}

type nodeEntry struct {
	*PulseEntry
	PulseProof *PulseProof
	Node       core.Node
}

func (ne *nodeEntry) hash(helper *merkleHelper) []byte {
	pulseHash := ne.PulseEntry.hash(helper)
	nodeInfoHash := helper.nodeInfoHash(pulseHash, ne.PulseProof.StateHash)
	return helper.nodeHash(ne.PulseProof.Signature.Bytes(), nodeInfoHash)
}

func nodeEntryByRole(pulseEntry *PulseEntry, nodeProofs map[core.Node]*PulseProof) map[core.StaticRole][]*nodeEntry {
	roleMap := make(map[core.StaticRole][]*nodeEntry)
	for node, pulseProof := range nodeProofs {
		role := node.Role()
		roleMap[role] = append(roleMap[role], &nodeEntry{
			PulseEntry: pulseEntry,
			Node:       node,
			PulseProof: pulseProof,
		})
	}
	return roleMap
}

func sortEntries(roleEntries []*nodeEntry) {
	sort.SliceStable(roleEntries, func(i, j int) bool {
		return roleEntries[i].Node.ID().Compare(roleEntries[j].Node.ID()) < 0
	})
}

func roleEntryRoot(roleEntries []*nodeEntry, helper *merkleHelper) ([]byte, error) {
	var roleEntriesHashes [][]byte
	for index, entry := range roleEntries {
		bucketEntryHash := helper.bucketEntryHash(uint32(index), entry.hash(helper))
		roleEntriesHashes = append(roleEntriesHashes, bucketEntryHash)
	}

	tree, err := treeFromHashList(roleEntriesHashes, helper.scheme.IntegrityHasher())
	if err != nil {
		return nil, errors.Wrap(err, "[ hash ] Failed to create tree")
	}

	return tree.Root(), nil
}
