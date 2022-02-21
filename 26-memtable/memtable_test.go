package memtable

import (
	"fmt"
	"testing"
)

func check(t *testing.T, err error, s string) {
	if err != nil {
		t.Fatalf("failure on %v\n%v", s, err)
	}
}

func get(t *testing.T, m Memtable, key string, expected_val string) {
	val, err := m.Get([]byte(key))
	check(t, err, fmt.Sprintf("get %v", key))
	if string(val) != expected_val {
		t.Fatalf("get %v: got %v; expected %v", key, string(val), expected_val)
	}
}

func has(t *testing.T, m Memtable, key string, expected_val bool) {
	val, err := m.Has([]byte(key))
	check(t, err, fmt.Sprintf("has %v", key))
	if val != expected_val {
		t.Fatalf("has %v: got %v; expected %v", key, val, expected_val)
	}
}

func put(t *testing.T, m *Memtable, key string, val string) {
	err := m.Put([]byte(key), []byte(val))
	check(t, err, fmt.Sprintf("put key %v val %v", key, val))
}

func delete(t *testing.T, m *Memtable, key string) {
	err := m.Delete([]byte(key))
	check(t, err, fmt.Sprintf("delete %v", key))
}

func TestMemtable(t *testing.T) {
	m := Memtable{}

	put(t, &m, "potato", "chip")
	put(t, &m, "tomato", "sauce")
	put(t, &m, "banana", "split")

	// get / has potato
	get(t, m, "potato", "chip")
	has(t, m, "potato", true)

	// overwrite potato
	put(t, &m, "potato", "salad")
	get(t, m, "potato", "salad")

	// delete potato
	delete(t, &m, "potato")
	has(t, m, "potato", false)

	// verify other pairs not changed
	get(t, m, "tomato", "sauce")
	get(t, m, "banana", "split")
}

func key(t *testing.T, iterator Iterator, expected string) {
	if string(iterator.Key()) != expected {
		t.Fatalf("key: expected %v, got %v", expected, string(iterator.Key()))
	}
}

func value(t *testing.T, iterator Iterator, expected string) {
	if string(iterator.Value()) != expected {
		t.Fatalf("value: expected %v, got %v", string(iterator.Value()), expected)
	}
}

func next(t *testing.T, iterator Iterator, expected bool) {
	if iterator.Next() != expected {
		t.Fatalf("next: expected %v, got %v", expected, !expected)
	}
}


// func TestIterator(t *testing.T) {
// 	m := Memtable{}

// 	put(t, &m, "potato", "chip")
// 	put(t, &m, "tomato", "sauce")
// 	put(t, &m, "banana", "split")
// 	put(t, &m, "orange", "juice")
// 	put(t, &m, "almond", "milk")
// 	put(t, &m, "carrot", "cake")
// 	put(t, &m, "garlic", "bread")
// 	put(t, &m, "cherry", "pie")

// 	fmt.Println(m)
// 	iterator, _ := m.RangeScan([]byte("carrot"), []byte("potato"))
// 	key(t, iterator, "carrot")
// 	value(t, iterator, "cake")

// 	next(t, iterator, true)
// 	key(t, iterator, "cherry")
// 	value(t, iterator, "pie")

// 	next(t, iterator, true)
// 	key(t, iterator, "garlic")
// 	value(t, iterator, "bread")

// 	next(t, iterator, true)
// 	key(t, iterator, "orange")
// 	value(t, iterator, "juice")

// 	next(t, iterator, false)
// 	key(t, iterator, "")
// 	value(t, iterator, "")
// }

// func TestIteratorEmpty(t *testing.T) {
// 	m := Memtable{}

// 	iterator, _ := m.RangeScan([]byte("carrot"), []byte("potato"))
// 	key(t, iterator, "")
// 	value(t, iterator, "")
// 	next(t, iterator, false)
// }

// func TestIteratorKeysNotInTable(t *testing.T) {
// 	m := Memtable{}

// 	iterator, _ := m.RangeScan([]byte("artichoke"), []byte(""))
// 	key(t, iterator, "")
// 	value(t, iterator, "")
// 	next(t, iterator, false)
// }
