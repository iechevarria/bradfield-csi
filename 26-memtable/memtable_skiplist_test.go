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
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(RandomLevel(8))
	// }
}

