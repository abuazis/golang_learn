package main

import "fmt"

func main() {
	//! Operator Aritmatika
	var value = ((( 2 + 6) % 3) * 4 - 2) / 3

	fmt.Printf("value: %d\n", value)

	//! Operator Perbandingan
	var result = ((( 2 + 6) % 3) * 4 - 2) / 3
	var isEqual = (result == 2)

	fmt.Printf("nilai: %d (%t) \n", value, isEqual)

	//! Operator Logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t(%t) \n", leftReverse)
}