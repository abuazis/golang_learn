package main

import (
	//! Pemanfaatan Alias Ketika Import Package
	f "fmt"

	//! Import Dengan Prefix Tanda Titik
	. "golang_learn/18_property_modifier/library"
)

func main() {
	SayHello("ethan")

	var s1 = Student{"ethan", 21}
    f.Println("name ", s1.Name)
    f.Println("grade", s1.Grade)

	sayHello2("ethan")

	f.Printf("Name  : %s\n", Student2.Name)
    f.Printf("Grade : %d\n", Student2.Grade)
}