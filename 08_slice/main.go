package main

import "fmt"

func main() {
	//! Inisialisasi Slice
	// fruits[x:y]
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])
	
	var _ = []string{"apple", "grape"} // slice
	var _ = [2]string{"banana", "melon"} // array
	var _ = [...]string{"papaya", "orange"} // array

	//! Hubungan Slice Dengan Array & Operasi Slice
	var fruitsA = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruitsA[0:2] // indeks ke-0, sampai sebelum indeks ke-2

	_ = fruitsA[0:4] // indeks ke-0, sampai sebelum indeks ke-4
	_ = fruitsA[0:0] // slice kosong, tidak ada elemen sebelum indeks ke-0
	_ = fruitsA[4:4] // slice kosong, tidak ada elemen yang dimulai indeks ke-4
	_ = fruitsA[:] // semua elemen
	_ = fruitsA[2:] // semua elemen mulai indeks ke-2
	_ = fruitsA[:2] // semua elemen sebelum indeks ke-2

	fmt.Println(newFruits)

	//! Slice Merupakan Tipe Data Reference
	var fruiters = []string{"apple", "grape", "banana", "melon"}

	var aFruiters = fruiters[0:3]
	var bFruiters = fruiters[1:4]

	var aaFruiters = aFruiters[1:2]
	var baFruiters = bFruiters[0:1]

	fmt.Println(fruiters) // [apple grape banana melon]
	fmt.Println(aFruiters) // [apple grape banana]
	fmt.Println(bFruiters) // [grape banana melon]
	fmt.Println(aaFruiters) // [grape]
	fmt.Println(baFruiters) // [grape]

	baFruiters[0] = "pineapple"

	fmt.Println(fruiters) // [apple pineapple banana melon]
	fmt.Println(aFruiters) // [apple pineapple banana]
	fmt.Println(bFruiters) // [pineapple banana melon]
	fmt.Println(aaFruiters) // [pineapple]
	fmt.Println(baFruiters) // [pineapple]

	//! Fungsi len()
	// Menghitung jumlah elemen slice yang ada
	
	var fruitsB = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruitsB))

	//! Fungsi cap()
	// Menghitung lebar atau kapasitas maksimum slice

	// [buah buah buah buah]
	var fruitsC = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruitsC)) // len: 4
	fmt.Println(cap(fruitsC)) // cap: 4

	// [buah buah buah ----]
	var aFruitsC = fruitsC[0:3]
	fmt.Println(len(aFruitsC)) // len: 3
	fmt.Println(cap(aFruitsC)) // cap: 4

	// ---- [buah buah buah]
	var bFruitsC = fruitsC[1:4]
	fmt.Println(len(bFruitsC)) // len: 3
	fmt.Println(cap(bFruitsC)) // cap: 3

	//! Fungsi append()
	// Menambahkan elemen pada slice

	var dFruits = []string{"apple", "grape", "banana"}
	var eFruits = append(dFruits, "papaya")

	fmt.Println(dFruits) // ["apple", "grape", "banana"]
	fmt.Println(eFruits) // ["apple", "grape", "banana", "papaya"]

	var fFruits = []string{"apple", "grape", "banana"}

	var hFruits = fFruits[0:2]

	fmt.Println(cap(hFruits)) // cap: 3
	fmt.Println(len(hFruits)) // len: 2

	fmt.Println(fFruits) // ["apple", "grape", "banana"]
	fmt.Println(hFruits) // ["apple", "grape"]

	var iFruits = append(hFruits, "manggo")

	fmt.Println(fFruits) // ["apple", "grape", "manggo"]
	fmt.Println(hFruits) // ["apple", "grape"]
	fmt.Println(iFruits) // ["apple", "grape", "manggo"]

	//! Fungsi copy()
	// Men-copy elements slice

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnapple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // [watermelon pinnapple apple]
	fmt.Println(src) // [watermelon pinnapple apple orange]
	fmt.Println(n) // 3

	dest := []string{"potato", "potato", "potato"}
	srce := []string{"watermelon", "pinnaple"}
	n2 := copy(dest, srce)

	fmt.Println(dest) // [potato pinnapple apple]
	fmt.Println(srce) // [watermelon pinnapple]
	fmt.Println(n2) // 2

	//! Pengaksesan Elemen Slice Dengan 3 Indeks
	var jFruits = []string{"apple", "grape", "banana"}
	var kFruits = jFruits[0:2]
	var lFruits = jFruits[0:2:2]

	fmt.Println(jFruits) // [apple grape banana]
	fmt.Println(len(jFruits)) // len: 3
	fmt.Println(cap(jFruits)) // cap: 3

	fmt.Println(kFruits) // [apple grape]
	fmt.Println(len(kFruits)) // len: 2
	fmt.Println(cap(kFruits)) // cap: 3

	fmt.Println(lFruits) // [apple grape]
	fmt.Println(len(lFruits)) // len: 2
	fmt.Println(cap(lFruits)) // cap: 2
}