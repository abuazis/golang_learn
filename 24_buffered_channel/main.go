package main

import (
	"fmt"
	"runtime"
)

func main() {
	//! Penerapan Buffered Channel
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2)

	go func() {
		for {
			// Pengiriman data
			i := <- messages
			fmt.Println("receive data", i)
		}
	}()
	
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		// Penerimaan data
		messages <- i
	}
}