package main

import "fmt"

func main() {
	//! Perulangan Menggunakan Keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("ANGKA", i)
	}

	//! Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	var i = 0

	for i < 5 {
		fmt.Println("ANGKA", i)
		i++
	}

	//! Penggunaan Keyword for Tanpa Argumen
	var x = 0

	for {
		fmt.Println("ANGKA", x)

		x++
		if x == 5 {
			break
		}
	}

	//! Penggunaan break & continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("ANGKA", i)
	}

	//! Perulangan Bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	//! Pemanfaatan Label Dalam Perulangan
	outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
