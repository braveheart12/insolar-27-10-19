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

package store

import (
	"context"
	"io"
	"sync"

	"github.com/dgraph-io/badger"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/pkg/errors"
)

// BadgerDB is a badger DB implementation.
type BadgerDB struct {
	backend   *badger.DB
	extraOpts BadgerOptions
}

type BadgerOptions struct {
	// ValueLogDiscardRatio set parameter for RunValueLogGC badger.DB method.
	valueLogDiscardRatio float64
}

type BadgerOption func(*BadgerOptions)

// ValueLogDiscardRatio configures values files garbage collection discard ratio.
// If value is greater than zero, NewBadgerDB starts values garbage collection in detached goroutine.
//
// More info about how it works in documentation of badger.DB's RunValueLogGC method.
func ValueLogDiscardRatio(value float64) BadgerOption {
	return func(s *BadgerOptions) {
		s.valueLogDiscardRatio = value
	}
}

// NewBadgerDB creates new BadgerDB instance.
// Creates new badger.DB instance with provided working dir and use it as backend for BadgerDB.
func NewBadgerDB(opts badger.Options, extras ...BadgerOption) (*BadgerDB, error) {
	b := &BadgerDB{}
	for _, opt := range extras {
		opt(&b.extraOpts)
	}

	// always allow to truncate vlog if necessary (actually it should have been a default behavior)
	opts.Truncate = true
	bdb, err := badger.Open(opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open badger")
	}
	b.backend = bdb

	return b, nil
}

func (b *BadgerDB) Backend() *badger.DB {
	return b.backend
}

// RunValueGC run badger values garbage collection
// Now it has to be called only after pulse finalization to
// exclude running GC during process of backup-replication
func (b *BadgerDB) RunValueGC(ctx context.Context) {
	if b.extraOpts.valueLogDiscardRatio > 0 {
		logger := inslogger.FromContext(ctx)
		logger.Info("BadgerDB: values GC start")
		defer logger.Info("BadgerDB: values GC end")

		err := b.backend.RunValueLogGC(b.extraOpts.valueLogDiscardRatio)
		if err != nil && err != badger.ErrNoRewrite {
			logger.Errorf("BadgerDB: GC failed with error: %v", err.Error())
		}
	}
}

// Stop gracefully stops all disk writes. After calling this, it's safe to kill the process without losing data.
func (b *BadgerDB) Stop(ctx context.Context) error {
	logger := inslogger.FromContext(ctx)
	defer logger.Info("BadgerDB: database closed")

	logger.Info("BadgerDB: closing database...")
	return b.backend.Close()
}

// Get returns value for specified key or an error. A copy of a value will be returned (i.e. getting large value can be
// long).
func (b *BadgerDB) Get(key Key) (value []byte, err error) {
	fullKey := append(key.Scope().Bytes(), key.ID()...)

	err = b.backend.View(func(txn *badger.Txn) error {
		item, err := txn.Get(fullKey)
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})

	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return
}

// Set stores value for a key.
func (b *BadgerDB) Set(key Key, value []byte) error {
	fullKey := append(key.Scope().Bytes(), key.ID()...)

	err := b.backend.Update(func(txn *badger.Txn) error {
		return txn.Set(fullKey, value)
	})

	return err
}

// Delete deletes value for a key.
func (b *BadgerDB) Delete(key Key) error {
	fullKey := append(key.Scope().Bytes(), key.ID()...)
	err := b.backend.Update(func(txn *badger.Txn) error {
		return txn.Delete(fullKey)
	})

	return err
}

// Backup creates backup.
func (b *BadgerDB) Backup(w io.Writer, since uint64) (uint64, error) {
	return b.backend.Backup(w, since)
}

// NewReadIterator returns new Iterator over the store.
func NewReadIterator(db *badger.DB, pivot Key, reverse bool) Iterator {
	bi := badgerIterator{pivot: pivot, reverse: reverse}
	bi.txn = db.NewTransaction(false)
	opts := badger.DefaultIteratorOptions
	opts.Reverse = reverse
	bi.it = bi.txn.NewIterator(opts)
	return &bi
}

// NewIterator returns new Iterator over the store.
func (b *BadgerDB) NewIterator(pivot Key, reverse bool) Iterator {
	bi := badgerIterator{pivot: pivot, reverse: reverse}
	bi.txn = b.backend.NewTransaction(false)
	opts := badger.DefaultIteratorOptions
	opts.Reverse = reverse
	bi.it = bi.txn.NewIterator(opts)
	return &bi
}

type badgerIterator struct {
	once      sync.Once
	pivot     Key
	reverse   bool
	txn       *badger.Txn
	it        *badger.Iterator
	prevKey   []byte
	prevValue []byte
}

func (bi *badgerIterator) Close() {
	bi.it.Close()
	bi.txn.Discard()
}

func (bi *badgerIterator) Next() bool {
	scope := bi.pivot.Scope().Bytes()
	bi.once.Do(func() {
		prefix := append(bi.pivot.Scope().Bytes(), bi.pivot.ID()...)
		bi.it.Seek(prefix)
	})
	if !bi.it.ValidForPrefix(scope) {
		return false
	}

	k := bi.it.Item().KeyCopy(nil)
	bi.prevKey = k[len(scope):]
	v, err := bi.it.Item().ValueCopy(nil)
	if err != nil {
		return false
	}
	bi.prevValue = v

	bi.it.Next()
	return true
}

func (bi *badgerIterator) Key() []byte {
	return bi.prevKey
}

func (bi *badgerIterator) Value() ([]byte, error) {
	return bi.prevValue, nil
}
