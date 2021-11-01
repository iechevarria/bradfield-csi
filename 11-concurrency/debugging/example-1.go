package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// need to actually pass into the function
		go func(x int) {
			fmt.Printf("launched goroutine %d\n", x)
		}(i)
	}
	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
