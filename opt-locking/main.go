package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var value int32 = 0

	// Perform a compare-and-swap operation
	previous := atomic.LoadInt32(&value) // Load the current value
	newValue := previous + 1             // Increment the value
	swapped := atomic.CompareAndSwapInt32(&value, previous, newValue)

	if swapped {
		fmt.Println("CAS succeeded")
	} else {
		fmt.Println("CAS failed")
	}

	fmt.Println("New value:", atomic.LoadInt32(&value))
}
