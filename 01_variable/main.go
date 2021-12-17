package main 

import "fmt"

func main() {
	//! Deklarasi Beserta Tipe Data (manifest typing)
	var firstName string = "John"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	//! Deklarasi Menggunakan Keyword var
	// var <nama-variabel> <tipe-data> 
	// var <nama-variabel> <tipe-data> = <nilai>

	var lastJob string
	var firstJob string = "farmer"

	fmt.Printf("halo farmer!\n")
	fmt.Printf("halo %s %s!\n", firstJob, lastJob)
	fmt.Println("halo", firstJob, lastName + "!")

	//! Deklarasi Tanpa Tipe Data (type inference)
	// Tipe data yang ditentukan dari nilainya

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstTask string = "Eat"
	
	// tanpa var, tanpa tipe data, menggunakan ":="
	lastTask := "Drink"
	lastTask = "Write"
	lastTask = "Wash"
	
	fmt.Printf("halo %s %s !\n", firstTask, lastTask)

	//! Deklarasi Multi Variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Printf("%s %s %s\n", first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Printf("%s %s %s\n", fourth, fifth, sixth)
	
	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	//! Varaibel Underscore _
	_ = "belajar mudah"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"

	fmt.Printf("%s\n", name);

	//! Deklarasi Variabel Menggunakan Keyword new
	job := new(string)

	fmt.Println(job)
	fmt.Println(*job)
}