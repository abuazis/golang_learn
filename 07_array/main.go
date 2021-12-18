package main

import "fmt"

func main() {
	//! Pengisian Elemen Array Melebihi Nilai Alokasi Awal
	var names [4]string

	names[0] = "zero"
	names[1] = "ichi"
	names[2] = "ni"
	names[3] = "san"
	// names[4] = "yon" // baris ini akan error

	fmt.Println(names[0], names[1], names[2], names[3])

	//! Inisialisasi Nilai Awal Array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	//! Inisialisasi Nilai Array Dengan Gaya Vertikal
	var things [4]string

	// cara horizontal
	things = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	things = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Isi semua element \t", things)

	//! Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var numbers = [...]int{2, 3, 4, 5, 6}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen  \t:", len(numbers))

	//! Array Multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//! Perulangan Elemen Array Menggunakan Keyword for
	var foods = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(foods); i++ {
		fmt.Printf("elemen %d : %s\n", i, foods[i])
	}

	//! Perulangan Elemen Array Menggunakan Keyword for - range
	var drinks = [4]string{"apple", "grape", "banana", "melon"}

	for i, drink := range drinks {
		fmt.Printf("elemen %d : %s\n", i, drink)
	}

	//! Penggunaan Variabel Underscore _ Dalam for - range
	var items = [4]string{"apple", "grape", "banana", "melon"}

	for _, item := range items {
		fmt.Printf("elemen %s\n", item)
	}

	// for i, _ := range fruits { }
	// atau
	// for i := range fruits { }

	//! Alokasi Elemen Array Menggunakan Keyword make
	var juices = make([]string, 2)

	juices[0] = "apple"
	juices[1] = "manggo"

	fmt.Println(juices)
}
