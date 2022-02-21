package memtable

import (
	"fmt"
	"testing"
)

func TestMemtable(t *testing.T) {
	m := NewMemtable(8)
	m.Put([]byte("potato"), []byte("salad"))
	m.Put([]byte("potato"), []byte("chip"))
	m.Put([]byte("jumbo"), []byte("shrimp"))
	m.Put([]byte("zebra"), []byte("stripe"))
	m.Put([]byte("zygote"), []byte("ok"))
	m.PrettyPrint()
	fmt.Println()
	m.Put([]byte("butter"), []byte("on toast"))
	m.PrettyPrint()
	fmt.Println(m.Get([]byte("potato")))
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(RandomLevel(8))
	// }
}

func TestMemtableAgain(t *testing.T) {
	m := NewMemtable(8)

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
		m.Put([]byte(w), []byte(fmt.Sprint(w, i)))
	}
	for _, w := range words {
		_, err := m.Get([]byte(w))
		if err != nil {
			t.Fatalf("got an error")
		}
	}

	m.PrettyPrint()
}
