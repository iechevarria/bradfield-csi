package main

import (
	"fmt"
	"unsafe"
)

func SpecialLen(s string) int {
	// Given a string, return its length without using len
	p := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(1))
	return *(*int)(p)
}

func SpecialPoint(p point) int {
	// TODO
}

func main () {
	x := "HELLO WORLD"
	fmt.Println(SpecialLen(x))
}
