package main

import (
	"fmt"
	"strings"
)

func main() {
	//! Penggunaan interface{}
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var _ map[string]interface{}

	_ = map[string]interface{}{
		"name"	    : "ethan hunt",
		"grade"		: 2,
		"breakfast" : []string{"apple", "manggo", "banana"},
	}

	//! Casting Variabel Interface Kosong
	var keep interface{}

	keep = 2
	var number = keep.(int) * 10
	fmt.Println(keep, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	//! Casting Variabel Interface Kosong Ke Objek Pointer
	type person struct {
		name string
		age	 int
	}

	var bio interface{} = &person{name: "wick", age: 27}
	var name = bio.(*person).name
	fmt.Println(name)

	//! Kombinasi Slice, map dan interface{}
	var person2 = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person2 {
		fmt.Println(each["name"], "age is", each["age"])
	}

	//! Dengan memanfaatkan slice dan interface{}, kita 
	//! bisa membuat data array yang isinya adalah bisa apa saja.
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}