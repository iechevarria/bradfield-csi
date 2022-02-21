package memtable

import (
	"fmt"
	"math"
	"math/rand"
)

const rate = 1.0 / math.E
const active = 0
const deleted = 1

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

type Memtable struct {
	MaxLevel int
	Size     int
	Header   Node
}

func NewMemtable(maxLevel int) *Memtable {
	header := &Node{Key: "", Value: nil, Level: 0, Forward: make([]*Node, maxLevel)}
	return &Memtable{MaxLevel: maxLevel, Header: *header}
}

type Node struct {
	Level   int
	Flags   uint8 // tombstone stored here. tbd if anything else
	Key     string
	Value   []byte
	Forward []*Node
}

func NewNode(key string, value []byte, level int) *Node {
	return &Node{
		Level:   level,
		Flags:   active,
		Key:     key,
		Value:   value,
		Forward: make([]*Node, level+1),
	}
}

func (m Memtable) RandomLevel() int {
	lvl := 0
	for rand.Float64() < rate && lvl < m.MaxLevel-1 {
		lvl += 1
	}
	return lvl
}

func (m Memtable) PrettyPrint() {
	fmt.Printf("%+v\n", m)
	cur := m.Header
	for {
		// gotta be a cleaner way to make this check
		if cur.Forward[0] == nil {
			return
		}
		addr := cur.Forward[0]
		cur = *cur.Forward[0]
		fmt.Printf("%p %+v\n", addr, cur)
	}
}

func (m *Memtable) Get(key []byte) ([]byte, error) {
	curNode := m.Header
	for i := m.Header.Level; i >= 0; i-- {
		// is there a less gross way to do this?
		for curNode.Forward[i] != nil && (*curNode.Forward[i]).Key <= string(key) {
			curNode = *curNode.Forward[i]
			if curNode.Key == string(key) && curNode.Flags != deleted {
				return curNode.Value, nil
			}
		}
	}

	return nil, KeyNotFound(key)
}

func (m *Memtable) Put(key []byte, value []byte) error {
	update := make([]*Node, m.MaxLevel)
	curNodeAddr := &m.Header
	for i := m.Header.Level; i >= 0; i-- {
		// is there a less gross way to do this?
		for (*curNodeAddr).Forward[i] != nil && (*(*curNodeAddr).Forward[i]).Key < string(key) {
			curNodeAddr = curNodeAddr.Forward[i]
		}
		update[i] = curNodeAddr
	}
	// is there a less gross way to do this?
	if (*curNodeAddr).Forward[0] != nil && (*(*curNodeAddr).Forward[0]).Key == string(key) {
		(*(*curNodeAddr).Forward[0]).Value = value
		(*(*curNodeAddr).Forward[0]).Flags = active
	} else {
		level := m.RandomLevel()
		if level > m.Header.Level {
			for i := m.Header.Level + 1; i <= level; i++ {
				update[i] = &m.Header
			}
			m.Header.Level = level
		}
		newNode := NewNode(string(key), value, level)
		for i := 0; i <= level; i++ {
			newNode.Forward[i] = (*update[i]).Forward[i]
			(*update[i]).Forward[i] = newNode
		}
		m.Size += 1
	}

	return nil
}

// func printUpdate(update []*Node) {
// 	fmt.Println("\n\n\n\n", update)
// 	for _, u := range update {
// 		if u != nil {
// 			fmt.Println(*u)
// 		}
// 	}
// }
