package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 = student{"john wick", 2}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)

	var s2 = student{"john wick", 2}
	fmt.Println("s2 before", s2.name)
    // john wick

    s2.changeName1("jason bourne")
    fmt.Println("s2 after changeName1", s2.name)
    // john wick

    s2.changeName2("ethan hunt")
    fmt.Println("s2 after changeName2", s2.name)
    // ethan hunt

	// pengaksesan method dari variabel objek biasa
	var s3 = student{"john wick", 21}
	s3.sayHello()

	// pengaksesan method dari variabel objek pointer
	var s4 = &student{"ethan hunt", 22}
	s4.sayHello()
}

//! Penerapan Method
type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i - 1]
}

//! Method Pointer
func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}