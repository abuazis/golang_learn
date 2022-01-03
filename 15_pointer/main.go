package main

import "fmt"

func main() {
	//! Penerapan Pointer
	var numberA int = 10
	var numberB *int = &numberA

	fmt.Println("numberA (value)	:", numberA)
	fmt.Println("numberA (value)	:", &numberA)

	fmt.Println("numberB (value)	:", *numberB)
	fmt.Println("numberB (value)	:", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value)	:", numberA)
	fmt.Println("numberA (value)	:", &numberA)

	fmt.Println("numberB (value)	:", *numberB)
	fmt.Println("numberB (value)	:", numberB)

	var number = 4
	fmt.Println("before :", number)

	change(&number, 10)
	fmt.Println("after :", number)
}

//! Parameter Pointer
func change(original *int, value int) {
	*original = value
}
