package main

import (
	"fmt"
	"unsafe"
)

type Point struct { 
	x int
	y int
}

func SpecialLen(s string) int {
	// Given a string, return its length without using len
	p := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(1))
	return *(*int)(p)
}

func SpecialPoint(p Point) int {
	// Given a Point struct, return its y coordinate without using p.y
	p_ := unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(1))
	return *(*int)(p_)
}

func main () {
	x := "HELLO WORLD"
	fmt.Println(SpecialLen(x))

	y := Point{55, 77}
	fmt.Println(SpecialPoint(y))
}
