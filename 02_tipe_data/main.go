package main

import "fmt"

func main() {
	//! Tipe Data Numerik Non Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1234123644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	//! Tipe Data Numerik Desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//! Tipe Data bool (Boolean)
	var exist bool = true

	fmt.Printf("exists? %t\n", exist)

	//! Tipe Data string
	var message string = "halo"
	fmt.Printf("message: %s\n", message)

	var messages = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang"`
	fmt.Printf("messages: %s\n", messages)
}