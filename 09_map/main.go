package main

import "fmt"

func main() {
	//! Penggunaan Map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	//! Penggunaan Nilai Map
	var data map[string]int
	// data["one"] = 1
	// akan muncul error

	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara horizontal
	var _ = map[string]int{"january": 50, "februari": 40}

	// cara vertical
	var _ = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	var _ = map[string]int{}
	var _ = make(map[string]int)
	var _ = *new(map[string]int)

	//! Iterasi Item Map Menggunakan for - range
	var chicken1 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken1 {
		fmt.Println(key, " \t:", val)
	}

	//! Menghapus Item Map
	var chicken2 = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken2))
	fmt.Println(chicken2)

	delete(chicken2, "januari")

	fmt.Println(len(chicken2))
	fmt.Println(chicken2)

	//! Deteksi Keberadaan Item Dengan Key Tertentu
	var chicken3 = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken3["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	//! Kombinasi Slice & Map
	var chickens = []map[string]string{
		{"name": "chicken blue", 	"gender": "male"},
		{"name": "chicken red", 	"gender": "male"},
		{"name": "chicken yellow", 	"gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	var _ = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "K001"},
		{"community": "chicken lovers"},
	}
}
