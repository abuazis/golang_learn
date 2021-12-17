package main

import "fmt"

func main() {
	//! Penggunaan Konstanta
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	const lastName string = "wick"
	fmt.Print("nice to meet you ", lastName)
}