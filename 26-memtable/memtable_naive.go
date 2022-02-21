package memtable

import (
	"fmt"
	"sort"
)

type DB interface {
	// Get gets the value for the given key. It returns an error if the
	// DB does not contain the key.
	Get(key []byte) (value []byte, err error)

	// Has returns true if the DB contains the given key.
	Has(key []byte) (ret bool, err error)

	// Put sets the value for the given key. It overwrites any previous value
	// for that key; a DB is not a multi-map.
	Put(key, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key []byte) error

	// RangeScan returns an Iterator (see below) for scanning through all
	// key-value pairs in the given range, ordered by key ascending.
	RangeScan(start, limit []byte) (Iterator, error)
}

type Iterator interface {
	// Next moves the iterator to the next key/value pair.
	// It returns false if the iterator is exhausted.
	Next() bool

	// Error returns any accumulated error. Exhausting all the key/value pairs
	// is not considered to be an error.
	Error() error

	// Key returns the key of the current key/value pair, or nil if done.
	Key() []byte

	// Value returns the value of the current key/value pair, or nil if done.
	Value() []byte
}

type KeyNotFound string

func (e KeyNotFound) Error() string {
	return fmt.Sprintf("Key not found")
}

type Entry struct {
	Key   string
	Value *[]byte
}

type Memtable struct {
	Entries []Entry
}

func (m Memtable) Get(key []byte) ([]byte, error) {
	for _, e := range m.Entries {
		if e.Key == string(key) {
			return *e.Value, nil
		}
	}
	return nil, KeyNotFound(key)
}

func (m Memtable) Has(key []byte) (bool, error) {
	_, err := m.Get(key)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (m *Memtable) Put(key []byte, value []byte) error {
	already_exists := false
	for _, e := range m.Entries {
		if e.Key == string(key) {
			*e.Value = value
			already_exists = true
		}
	}
	if !already_exists {
		m.Entries = append(m.Entries, Entry{string(key), &value})
	}
	sort.Slice(m.Entries, func(i, j int) bool {
		return m.Entries[i].Key < m.Entries[j].Key
	})
	return nil
}

func (m *Memtable) Delete(key []byte) error {
	for i, e := range m.Entries {
		if e.Key == string(key) {
			m.Entries[i] = m.Entries[len(m.Entries)-1]
			m.Entries = m.Entries[:len(m.Entries)-1]
			return nil
		}
	}
	return KeyNotFound(key)
}

func clamp(num int) int {
    if num < 0 {
        return 0
    }
    return num
}

func (m Memtable) RangeScan(start, limit []byte) (Iterator, error) {
	for i, e := range m.Entries {
		if e.Key > string(start) {
			return &RangeIterator{Err: nil, Index: clamp(i - 1), Limit: string(limit), Table: m}, nil
		}
	}
	return &RangeIterator{Err: nil, Index: len(m.Entries), Limit: string(limit), Table: m}, nil
}

type RangeIterator struct {
	Err   error
	Index int
	Limit string
	Table Memtable
}

func (r *RangeIterator) Next() bool {
    // already at the end of the table
    if r.Index >= len(r.Table.Entries) {
        return false
    }

    // just reached limit
	if r.Table.Entries[r.Index + 1].Key >= r.Limit {
        r.Index = len(r.Table.Entries)
		return false
	}

	r.Index += 1
	return true
}

func (r *RangeIterator) Error() error {
	return r.Err
}

func (r *RangeIterator) Key() []byte {
	if r.Index >= len(r.Table.Entries) {
		return nil
	}
	return []byte(r.Table.Entries[r.Index].Key)
}

func (r *RangeIterator) Value() []byte {
	if r.Index >= len(r.Table.Entries) {
		return nil
	}
	return *r.Table.Entries[r.Index].Value
}
