package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContains0 = filter(data, func (each string) bool {
		return strings.Contains(each, "o")
	})
	var dataContains5 = filter(data, func (each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContains0)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataContains5)
	// filter jumlah angka "5" : [jason ethan]
}

//! Alias Skema Closure
type FilterCallback func(string) bool

//! Penerapan Fungsi Sebagai Parameter (Callback)
func filter(data []string, callback FilterCallback) []string {
	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}