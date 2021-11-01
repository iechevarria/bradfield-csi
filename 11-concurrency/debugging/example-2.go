package main

import (
	"fmt"
)

const numTasks = 3

func main() {
	// previously, this line was used:
	// var done chan struct{}
	// fix: need to initialize the channel
	done := make(chan struct{})
	for i := 0; i < numTasks; i++ {
		go func() {
			fmt.Println("running task...")

			// Signal that task is done
			done <- struct{}{}
		}()
	}

	// Wait for tasks to complete
	for i := 0; i < numTasks; i++ {
		<-done
	}
	fmt.Printf("all %d tasks done!\n", numTasks)
}
