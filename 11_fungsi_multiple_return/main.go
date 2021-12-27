package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran \t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran \t\t: %.2f \n", circumference)

	diameter = 7
	area, circumference = calculatePredefined(diameter)

	fmt.Printf("luas lingkaran \t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran \t\t: %.2f \n", circumference)
}

//! Penerapan Funsi Multiple Return
func calculate(d float64) (float64, float64) {
	// hitung luas 
	var area = math.Pi * math.Pow(d / 2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

//! Fungsi Dengan Predefined Return Value
func calculatePredefined(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi * d

	return
}