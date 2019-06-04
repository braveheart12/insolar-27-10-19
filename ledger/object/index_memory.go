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

package object

import (
	"context"
	"sort"
	"sync"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/pkg/errors"
	"go.opencensus.io/stats"
)

// filamentCache is a thread-safe wrapper around FilamentIndex struct, because FilamentIndex
// is a protobuf-generated struct.
// Also it stores some meta-info, that is required for the work process
type filamentCache struct {
	sync.RWMutex

	objectMeta  FilamentIndex
	pendingMeta pendingMeta
}

// pendingMeta contains info for calculating pending requests states
// The structure contains full chain of pendings (they are grouped by pulse).
// Groups are sorted by pulse, from a lowest to a highest
// There are a few maps inside, that help not to search through full fillament every SetRequest/SetResult
type pendingMeta struct {
	isStateCalculated bool
	fullFilament      []chainLink

	notClosedRequestsIds   []insolar.ID
	notClosedRequestsIndex map[insolar.PulseNumber]map[insolar.ID]*record.Request
	requestPNIndex         map[insolar.ID]insolar.PulseNumber
}

type chainLink struct {
	PN         insolar.PulseNumber
	RecordsIDs []insolar.ID
}

func (i *filamentCache) lifeline() (Lifeline, error) {
	i.RLock()
	defer i.RUnlock()

	return CloneIndex(i.objectMeta.Lifeline), nil
}

func (i *filamentCache) setLifeline(lifeline Lifeline, pn insolar.PulseNumber) {
	i.Lock()
	defer i.Unlock()

	i.objectMeta.Lifeline = lifeline
	i.objectMeta.LifelineLastUsed = pn
}

func (i *filamentCache) setLifelineLastUsed(pn insolar.PulseNumber) {
	i.Lock()
	defer i.Unlock()

	i.objectMeta.LifelineLastUsed = pn
}

// InMemoryIndex is a in-memory storage, that stores a collection of IndexBuckets
type InMemoryIndex struct {
	recordStorage RecordStorage

	bucketsLock sync.RWMutex
	buckets     map[insolar.PulseNumber]map[insolar.ID]*filamentCache
}

// NewInMemoryIndex creates a new InMemoryIndex
func NewInMemoryIndex(recordStorage RecordStorage) *InMemoryIndex {
	return &InMemoryIndex{
		recordStorage: recordStorage,
		buckets:       map[insolar.PulseNumber]map[insolar.ID]*filamentCache{},
	}
}

func (i *InMemoryIndex) createBucket(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) *filamentCache {
	i.bucketsLock.Lock()
	defer i.bucketsLock.Unlock()

	bucket := &filamentCache{
		objectMeta: FilamentIndex{
			ObjID:          objID,
			PendingRecords: []insolar.ID{},
		},
		pendingMeta: pendingMeta{
			fullFilament:           []chainLink{},
			notClosedRequestsIds:   []insolar.ID{},
			notClosedRequestsIndex: map[insolar.PulseNumber]map[insolar.ID]*record.Request{},
			requestPNIndex:         map[insolar.ID]insolar.PulseNumber{},
			isStateCalculated:      false,
		},
	}

	objsByPn, ok := i.buckets[pn]
	if !ok {
		objsByPn = map[insolar.ID]*filamentCache{}
		i.buckets[pn] = objsByPn
	}
	objsByPn[objID] = bucket

	inslogger.FromContext(ctx).Debugf("[createBucket] create bucket for obj - %v was created successfully", objID.DebugString())
	return bucket
}

func (i *InMemoryIndex) bucket(pn insolar.PulseNumber, objID insolar.ID) *filamentCache {
	i.bucketsLock.RLock()
	defer i.bucketsLock.RUnlock()

	objsByPn, ok := i.buckets[pn]
	if !ok {
		return nil
	}

	return objsByPn[objID]
}

// Set sets a lifeline to a bucket with provided pulseNumber and ID
func (i *InMemoryIndex) Set(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID, lifeline Lifeline) error {
	b := i.bucket(pn, objID)
	if b == nil {
		b = i.createBucket(ctx, pn, objID)
	}
	b.setLifeline(lifeline, pn)

	stats.Record(ctx,
		statBucketAddedCount.M(1),
	)

	inslogger.FromContext(ctx).Debugf("[Set] lifeline for obj - %v was set successfully", objID.DebugString())
	return nil
}

// SetBucket adds a bucket with provided pulseNumber and ID
func (i *InMemoryIndex) SetBucket(ctx context.Context, pn insolar.PulseNumber, bucket FilamentIndex) error {
	i.bucketsLock.Lock()
	defer i.bucketsLock.Unlock()

	bucks, ok := i.buckets[pn]
	if !ok {
		bucks = map[insolar.ID]*filamentCache{}
		i.buckets[pn] = bucks
	}

	bucks[bucket.ObjID] = &filamentCache{
		objectMeta: bucket,
		pendingMeta: pendingMeta{
			notClosedRequestsIds:   []insolar.ID{},
			fullFilament:           []chainLink{},
			isStateCalculated:      false,
			requestPNIndex:         map[insolar.ID]insolar.PulseNumber{},
			notClosedRequestsIndex: map[insolar.PulseNumber]map[insolar.ID]*record.Request{},
		},
	}

	stats.Record(ctx,
		statBucketAddedCount.M(1),
	)

	return nil
}

// ForID returns a lifeline from a bucket with provided PN and ObjID
func (i *InMemoryIndex) ForID(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) (Lifeline, error) {
	b := i.bucket(pn, objID)
	if b == nil {
		return Lifeline{}, ErrLifelineNotFound
	}
	return b.lifeline()
}

// ForPNAndJet returns a collection of buckets for a provided pn and jetID
func (i *InMemoryIndex) ForPNAndJet(ctx context.Context, pn insolar.PulseNumber, jetID insolar.JetID) []FilamentIndex {
	i.bucketsLock.Lock()
	defer i.bucketsLock.Unlock()

	bucks, ok := i.buckets[pn]
	if !ok {
		return nil
	}

	res := []FilamentIndex{}

	for _, b := range bucks {
		if b.objectMeta.Lifeline.JetID != jetID {
			continue
		}

		clonedLfl := CloneIndex(b.objectMeta.Lifeline)
		var clonedRecords []insolar.ID

		clonedRecords = append(clonedRecords, b.objectMeta.PendingRecords...)

		res = append(res, FilamentIndex{
			ObjID:            b.objectMeta.ObjID,
			Lifeline:         clonedLfl,
			LifelineLastUsed: b.objectMeta.LifelineLastUsed,
			PendingRecords:   clonedRecords,
		})
	}

	return res
}

// SetLifelineUsage updates a last usage fields of a bucket for a provided pulseNumber and an object id
func (i *InMemoryIndex) SetLifelineUsage(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) error {
	b := i.bucket(pn, objID)
	if b == nil {
		return ErrLifelineNotFound
	}

	b.setLifelineLastUsed(pn)

	return nil
}

// DeleteForPN deletes all buckets for a provided pulse number
func (i *InMemoryIndex) DeleteForPN(ctx context.Context, pn insolar.PulseNumber) {
	i.bucketsLock.Lock()
	defer i.bucketsLock.Unlock()

	bucks, ok := i.buckets[pn]
	if !ok {
		return
	}

	delete(i.buckets, pn)

	stats.Record(ctx,
		statBucketRemovedCount.M(int64(len(bucks))),
	)

	for _, buck := range bucks {
		stats.Record(ctx,
			statObjectPendingRequestsInMemoryRemovedCount.M(int64(len(buck.objectMeta.PendingRecords))),
		)
	}
}

// SetRequest sets a request for a specific object
func (i *InMemoryIndex) SetRequest(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID, reqID insolar.ID, req record.Request) error {
	b := i.bucket(pn, objID)
	if b == nil {
		return ErrLifelineNotFound
	}

	b.Lock()
	defer b.Unlock()

	b.objectMeta.PendingRecords = append(b.objectMeta.PendingRecords, reqID)

	isInserted := false
	for i, chainPart := range b.pendingMeta.fullFilament {
		if chainPart.PN == pn {
			b.pendingMeta.fullFilament[i].RecordsIDs = append(b.pendingMeta.fullFilament[i].RecordsIDs, reqID)
			isInserted = true
		}
	}

	if !isInserted {
		b.pendingMeta.fullFilament = append(b.pendingMeta.fullFilament, chainLink{RecordsIDs: []insolar.ID{reqID}, PN: pn})
		sort.Slice(b.pendingMeta.fullFilament, func(i, j int) bool {
			return b.pendingMeta.fullFilament[i].PN < b.pendingMeta.fullFilament[j].PN
		})
	}

	b.pendingMeta.requestPNIndex[reqID] = pn

	_, ok := b.pendingMeta.notClosedRequestsIndex[pn]
	if !ok {
		b.pendingMeta.notClosedRequestsIndex[pn] = map[insolar.ID]*record.Request{}
	}
	b.pendingMeta.notClosedRequestsIndex[pn][reqID] = &req
	b.pendingMeta.notClosedRequestsIds = append(b.pendingMeta.notClosedRequestsIds, reqID)

	stats.Record(ctx,
		statObjectPendingRequestsInMemoryAddedCount.M(int64(1)),
	)

	return nil

}

// SetResult sets a result for a specific object. Also, if there is a not closed request for a provided result,
// the request will be closed
func (i *InMemoryIndex) SetResult(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID, resID insolar.ID, res record.Result) error {
	b := i.bucket(pn, objID)
	if b == nil {
		return ErrLifelineNotFound
	}

	b.Lock()
	defer b.Unlock()

	b.objectMeta.PendingRecords = append(b.objectMeta.PendingRecords, resID)

	isInserted := false
	for i, chainPart := range b.pendingMeta.fullFilament {
		if chainPart.PN == pn {
			b.pendingMeta.fullFilament[i].RecordsIDs = append(b.pendingMeta.fullFilament[i].RecordsIDs, resID)
			isInserted = true
		}
	}

	if !isInserted {
		b.pendingMeta.fullFilament = append(b.pendingMeta.fullFilament, chainLink{RecordsIDs: []insolar.ID{resID}, PN: pn})
		sort.Slice(b.pendingMeta.fullFilament, func(i, j int) bool {
			return b.pendingMeta.fullFilament[i].PN < b.pendingMeta.fullFilament[j].PN
		})
	}

	reqPN, ok := b.pendingMeta.requestPNIndex[*res.Request.Record()]
	if ok {
		delete(b.pendingMeta.notClosedRequestsIndex[reqPN], *res.Request.Record())
		for i := 0; i < len(b.pendingMeta.notClosedRequestsIds); i++ {
			if b.pendingMeta.notClosedRequestsIds[i] == *res.Request.Record() {
				b.pendingMeta.notClosedRequestsIds = append(b.pendingMeta.notClosedRequestsIds[:i], b.pendingMeta.notClosedRequestsIds[i+1:]...)
				break
			}
		}
	}

	stats.Record(ctx,
		statObjectPendingRequestsInMemoryAddedCount.M(int64(1)),
	)

	return nil
}

// SetFilament adds a slice of records to an object with provided id and pulse. It's assumed, that the method is
// called for setting records from another light, during the process of filling full chaing of pendings
func (i *InMemoryIndex) SetFilament(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID, filPN insolar.PulseNumber, recs []record.MaterialWithId) error {
	b := i.bucket(pn, objID)
	if b == nil {
		return ErrLifelineNotFound
	}

	b.Lock()
	defer b.Unlock()

	recsIds := make([]insolar.ID, len(recs))
	for idx, rec := range recs {
		recsIds[idx] = rec.Id
		err := i.recordStorage.Set(ctx, rec.Id, rec.Record)
		if err != nil {
			return errors.Wrap(err, "filament update failed")
		}
	}

	b.pendingMeta.fullFilament = append(b.pendingMeta.fullFilament, chainLink{RecordsIDs: recsIds, PN: filPN})
	sort.Slice(b.pendingMeta.fullFilament, func(i, j int) bool {
		return b.pendingMeta.fullFilament[i].PN < b.pendingMeta.fullFilament[j].PN
	})

	return nil
}

// RefreshState recalculates state of the chain, marks requests as closed and opened.
func (i *InMemoryIndex) RefreshState(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) error {
	b := i.bucket(pn, objID)
	if b == nil {
		return ErrLifelineNotFound
	}

	b.Lock()
	defer b.Unlock()

	for _, chainLink := range b.pendingMeta.fullFilament {
		for _, recID := range chainLink.RecordsIDs {
			rec, err := i.recordStorage.ForID(ctx, recID)
			if err != nil {
				return errors.Wrap(err, "failed to refresh an index state")
			}
			switch r := record.Unwrap(rec.Virtual).(type) {
			case *record.Request:
				b.pendingMeta.notClosedRequestsIndex[chainLink.PN][*r.Object.Record()] = r
				b.pendingMeta.requestPNIndex[*r.Object.Record()] = chainLink.PN
			case *record.Result:
				reqPN, ok := b.pendingMeta.requestPNIndex[*r.Request.Record()]
				if ok {
					delete(b.pendingMeta.notClosedRequestsIndex[reqPN], *r.Request.Record())
				}
			}
		}
	}

	isEarliestFound := false

	for i, chainLink := range b.pendingMeta.fullFilament {
		if len(b.pendingMeta.notClosedRequestsIndex[chainLink.PN]) != 0 {
			if !isEarliestFound {
				b.objectMeta.Lifeline.EarliestOpenRequest = b.pendingMeta.fullFilament[i].PN
				isEarliestFound = true
			}

			for _, ncr := range b.pendingMeta.notClosedRequestsIndex[chainLink.PN] {
				panic("refactor here")
				// b.pendingMeta.notClosedRequestsIds = append(b.pendingMeta.notClosedRequestsIds, *ncr)
			}
		}
	}

	b.pendingMeta.isStateCalculated = true

	return nil
}

// IsStateCalculated returns status of a pending filament. Was it calculated or not
func (i *InMemoryIndex) IsStateCalculated(ctx context.Context, currentPN insolar.PulseNumber, objID insolar.ID) (bool, error) {
	b := i.bucket(currentPN, objID)
	if b == nil {
		return false, ErrLifelineNotFound
	}

	b.RLock()
	defer b.RUnlock()

	return b.pendingMeta.isStateCalculated, nil
}

// OpenRequestsForObjID returns open requests for a specific object
func (i *InMemoryIndex) OpenRequestsForObjID(ctx context.Context, currentPN insolar.PulseNumber, objID insolar.ID, count int) ([]record.Request, error) {
	b := i.bucket(currentPN, objID)
	if b == nil {
		return nil, ErrLifelineNotFound
	}

	b.RLock()
	defer b.RUnlock()

	if len(b.pendingMeta.notClosedRequestsIds) < count {
		count = len(b.pendingMeta.notClosedRequestsIds)
	}

	res := make([]record.Request, count)

	for idx := 0; idx < count; idx++ {
		rec, err := i.recordStorage.ForID(ctx, b.pendingMeta.notClosedRequestsIds[idx])
		if err != nil {
			return nil, err
		}

		switch r := record.Unwrap(rec.Virtual).(type) {
		case *record.Request:
			res[idx] = *r
		default:
			panic("filament is totally broken")
		}
	}

	return res, nil
}

// Records returns all the records for a provided object
func (i *InMemoryIndex) Records(ctx context.Context, currentPN insolar.PulseNumber, objID insolar.ID) ([]record.MaterialWithId, error) {
	b := i.bucket(currentPN, objID)
	if b == nil {
		return nil, ErrLifelineNotFound
	}

	b.RLock()
	defer b.RLock()

	res := make([]record.MaterialWithId, len(b.objectMeta.PendingRecords))
	for idx, id := range b.objectMeta.PendingRecords {
		rec, err := i.recordStorage.ForID(ctx, id)
		if err != nil {
			return nil, err
		}
		res[idx] = record.MaterialWithId{
			Record: rec,
			Id:     id,
		}
	}

	return res, nil
}
