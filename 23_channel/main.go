package main

import (
	"fmt"
	"runtime"
)

func main() {
	//! Penerapan Channel
	runtime.GOMAXPROCS(2)

	// Deklarasi channel
	var messages = make(chan string)

	var sayToHello = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		// Channel sebagai penerima
		messages <- data
	}

	go sayToHello("john wick")
	go sayToHello("ethan hunt")
	go sayToHello("jason bourne")

	// Channel sebagai pengirim
	var message1 = <-messages
	fmt.Println(message1)

	// Channel sebagai pengirim
	var message2 = <-messages
	fmt.Println(message2)

	// Channel sebagai pengirim
	var message3 = <-messages
	fmt.Println(message3)

	// Implementasi printMessage
	runtime.GOMAXPROCS(2)

    var messages2 = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages2 <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages2)
	}
}

//! Channel Sebagai Tipe Data Parameter
func printMessage(what chan string) {
	fmt.Println(<-what)
}