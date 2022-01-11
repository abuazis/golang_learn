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
}