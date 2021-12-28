package main

import "fmt"

func main() {
	//! Closure Disimpan Dalam Variabel (Anonymous Function)
	var getMinMax = func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0 :
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)

	//! Penggunaan Template String %v
	// menampilkan segala jenis data (dinamis)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	//! Immediately-Invoked Function Expression (IIFE)
	// Closure jenis ini dieksekusi langsung pada saat deklarasinya
	var newNumbers = func (min int) []int {
		var r []int

		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}

		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)

	// Jalankan closure nilai kembalian
	max = 3
	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)
}


//! Closure Sebagai Nilai Kembalian
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int

	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}