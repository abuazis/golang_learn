package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	} 
}

func main() {
	// Menentukan jumlah core / processor 
	// yang digunakan dalam eksekusi program
	runtime.GOMAXPROCS(2)

	// menjalankan dengan goroutine
	go print(5, "hello")
	// menjalankan tanpa goroutine
	print(5, "apa kabar")

	var input string
	// meng-capture semua karakter 
	// sebelum user menekan tombol enter
	fmt.Scanln(&input)
}