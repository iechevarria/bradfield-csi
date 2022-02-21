package crapdb

import (
	"fmt"
	"math"
	"math/rand"
)

const rate = 1.0 / math.E
const active = 0
const deleted = 1

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
	Deleted bool
	Key     string
	Value   []byte
	Forward []*Node
}

func NewNode(key string, value []byte, level int) *Node {
	return &Node{
		Level:   level,
		Deleted: false,
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
		for curNode.Forward[i] != nil && curNode.Forward[i].Key <= string(key) {
			curNode = *curNode.Forward[i]
			if curNode.Key == string(key) && !curNode.Deleted {
				return curNode.Value, nil
			}
		}
	}

	return nil, KeyNotFound(key)
}

func (m *Memtable) Has(key []byte) (bool, error) {
	_, err := m.Get(key)
	return err == nil, nil
}

func (m *Memtable) Delete(key []byte) error {
	curNodeAddr := &m.Header
	for i := m.Header.Level; i >= 0; i-- {
		for curNodeAddr.Forward[i] != nil && curNodeAddr.Forward[i].Key <= string(key) {
			curNodeAddr = curNodeAddr.Forward[i]
			if curNodeAddr.Key == string(key) {
				curNodeAddr.Deleted = true
				return nil
			}
		}
	}

	return KeyNotFound(key)
}

func (m *Memtable) Put(key []byte, value []byte) error {
	update := make([]*Node, m.MaxLevel)
	curNodeAddr := &m.Header
	for i := m.Header.Level; i >= 0; i-- {
		for curNodeAddr.Forward[i] != nil && curNodeAddr.Forward[i].Key < string(key) {
			curNodeAddr = curNodeAddr.Forward[i]
		}
		update[i] = curNodeAddr
	}
	if curNodeAddr.Forward[0] != nil && curNodeAddr.Forward[0].Key == string(key) {
		curNodeAddr.Forward[0].Value = value
		curNodeAddr.Forward[0].Deleted = false
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
			newNode.Forward[i] = update[i].Forward[i]
			update[i].Forward[i] = newNode
		}
		m.Size += 1
	}

	return nil
}

type RangeIterator struct {
	Err     error
	CurNode *Node
	Limit   string
}

func NewRangeIterator(curNode *Node, limit []byte) *RangeIterator {
	return &RangeIterator{
		Err:     nil,
		CurNode: curNode,
		Limit:   string(limit),
	}
}

func (m *Memtable) RangeScan(start, limit []byte) (Iterator, error) {
	if m.Size == 0 {
		return NewRangeIterator(nil, limit), nil
	}

	curNodeAddr := &m.Header
	for i := m.Header.Level; i >= 0; i-- {
		// is there a less gross way to do this?
		for curNodeAddr.Forward[i] != nil && curNodeAddr.Forward[i].Key < string(start) {
			curNodeAddr = curNodeAddr.Forward[i]
		}
	}

	curNodeAddr = curNodeAddr.Forward[0]
	return NewRangeIterator(curNodeAddr, limit), nil
}

func (r *RangeIterator) Next() bool {
	// terminated
	if r.CurNode == nil {
		return false
	}

	// continue until not on deleted node
	for ok := true; ok; ok = r.CurNode.Deleted {
		// at the end of the table
		if r.CurNode.Forward[0] == nil {
			r.CurNode = nil
			return false
		}
		// at limit
		if r.CurNode.Forward[0].Key >= r.Limit {
			r.CurNode = nil
			return false
		}

		// otherwise go forward
		r.CurNode = r.CurNode.Forward[0]
	}

	return true
}

func (r *RangeIterator) Error() error {
	return r.Err
}

func (r *RangeIterator) Key() []byte {
	if r.CurNode == nil {
		return nil
	}
	return []byte(r.CurNode.Key)
}

func (r *RangeIterator) Value() []byte {
	if r.CurNode == nil {
		return nil
	}
	return []byte(r.CurNode.Value)
}
