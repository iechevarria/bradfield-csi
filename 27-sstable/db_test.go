package crapdb

import (
	"fmt"
	"testing"
)

func check(t *testing.T, err error, s string) {
	if err != nil {
		t.Fatalf("failure on %v\n%v", s, err)
	}
}

func get(t *testing.T, db *CrapDB, key string, expected_val string) {
	val, err := db.Get([]byte(key))
	check(t, err, fmt.Sprintf("get %v", key))
	if string(val) != expected_val {
		t.Fatalf("get %v: got %v; expected %v", key, string(val), expected_val)
	}
}

func has(t *testing.T, db *CrapDB, key string, expected_val bool) {
	val, err := db.Has([]byte(key))
	check(t, err, fmt.Sprintf("has %v", key))
	if val != expected_val {
		t.Fatalf("has %v: got %v; expected %v", key, val, expected_val)
	}
}

func put(t *testing.T, db *CrapDB, key string, val string) {
	err := db.Put([]byte(key), []byte(val))
	check(t, err, fmt.Sprintf("put key %v val %v", key, val))
}

func delete(t *testing.T, db *CrapDB, key string) {
	err := db.Delete([]byte(key))
	check(t, err, fmt.Sprintf("delete %v", key))
}

func TestCrapDB1(t *testing.T) {
	db := NewCrapDB(8)

	put(t, db, "potato", "chip")
	put(t, db, "tomato", "sauce")
	put(t, db, "banana", "split")

	// get / has potato
	get(t, db, "potato", "chip")
	has(t, db, "potato", true)

	// overwrite potato
	put(t, db, "potato", "salad")
	get(t, db, "potato", "salad")

	// delete potato
	delete(t, db, "potato")
	has(t, db, "potato", false)

	// verify other pairs not changed
	get(t, db, "tomato", "sauce")
	get(t, db, "banana", "split")

	db.PrettyPrint()
}

func TestCrapDB2(t *testing.T) {
	db := NewCrapDB(8)
	put(t, db, "potato", "salad")
	put(t, db, "potato", "chip")
	put(t, db, "jumbo", "shrimp")
	put(t, db, "zebra", "stripe")
	put(t, db, "zumba", "class")
	put(t, db, "butter", "on toast")

	db.PrettyPrint()
}

func TestCrapDB3(t *testing.T) {
	db := NewCrapDB(8)

	words := []string{
		"words",
		"are",
		"cool",
		"what",
		"do",
		"you",
		"think",
		"look",
		"out",
		"we",
		"have",
		"even",
		"more",
		"cool",
		"words",
		"they",
		"are",
		"everywhere",
		"words",
		"like",
		"xylophone",
		"neat",
	}

	for i, w := range words {
		put(t, db, w, fmt.Sprint(w, i))
	}
	for _, w := range words {
		has(t, db, w, true)
	}

	db.PrettyPrint()
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

func TestIterator(t *testing.T) {
	db := NewCrapDB(8)

	put(t, db, "potato", "chip")
	put(t, db, "tomato", "sauce")
	put(t, db, "banana", "split")
	put(t, db, "orange", "juice")
	put(t, db, "almond", "milk")
	put(t, db, "carrot", "cake")
	put(t, db, "garlic", "bread")
	put(t, db, "cherry", "pie")
	db.PrettyPrint()

	iterator, _ := db.RangeScan([]byte("carrot"), []byte("potato"))
	key(t, iterator, "carrot")
	value(t, iterator, "cake")

	next(t, iterator, true)
	key(t, iterator, "cherry")
	value(t, iterator, "pie")

	next(t, iterator, true)
	key(t, iterator, "garlic")
	value(t, iterator, "bread")

	next(t, iterator, true)
	key(t, iterator, "orange")
	value(t, iterator, "juice")

	next(t, iterator, false)
	key(t, iterator, "")
	value(t, iterator, "")

	// handle multiple sequential deletions
	delete(t, db, "cherry")
	delete(t, db, "garlic")
	iterator, _ = db.RangeScan([]byte("carrot"), []byte("potato"))
	key(t, iterator, "carrot")
	value(t, iterator, "cake")

	next(t, iterator, true)
	key(t, iterator, "orange")
	value(t, iterator, "juice")

	next(t, iterator, false)
	key(t, iterator, "")
	value(t, iterator, "")

	// handle deletion at end
	delete(t, db, "orange")
	iterator, _ = db.RangeScan([]byte("carrot"), []byte("potato"))
	key(t, iterator, "carrot")
	value(t, iterator, "cake")

	next(t, iterator, false)
	key(t, iterator, "")
	value(t, iterator, "")

	db.PrettyPrint()
}

func TestIteratorEmpty(t *testing.T) {
	db := NewCrapDB(8)

	iterator, _ := db.RangeScan([]byte("carrot"), []byte("potato"))
	key(t, iterator, "")
	value(t, iterator, "")
	next(t, iterator, false)
}
