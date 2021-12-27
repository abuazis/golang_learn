package main

//! Import Banyak Package
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"john", "wick"}
	printMessage("halo", names)

	//! Penggunaan Fungsi rand.Seed()
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	divideByNumber(10, 2)
	divideByNumber(4, 0)
	divideByNumber(8, -4)
}

//! Penerapan Fungsi
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

//! Deklarasi Parameter Bertipe Data Sama
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min

	//! Fungsi Dengan Return Value / Nilai Balik
	return value
}

//! Penggunaan Keyword return Untuk Menghentikan Proses Dalam fungsi
func divideByNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d,\n", m, n, res)
}