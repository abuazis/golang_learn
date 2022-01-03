package main

import "fmt"

func main() {
	//! Penerapan Struct
	var s1 student
	s1.name = "John Wick"
	s1.grade = 2

	fmt.Println("name	:", s1.name)
	fmt.Println("grade	:", s1.grade)

	//! Inisialisasi Object Struct
	var s2 = student{}
	s2.name = "Wick"
	s2.grade = 2

	var s3 = student{"Ethan", 3}

	var s4 = student{name: "Jason", grade: 2}

	s4 = student{grade: 2, name: "Jason"}

	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)
	fmt.Println("student 4 :", s4.name)

	//! Variabel Objek Pointer
	var s5 *student = &s4
	fmt.Println("student 4, name :", s4.name)
	fmt.Println("student 5, name :", s5.name)

	s5.name = "ethan"
	fmt.Println("student 4, name :", s4.name)
	fmt.Println("student 5, name :", s5.name)

	//! Implementasi Embedded Struct
	var s6 = scientist{}
	s6.name = "wick"
	s6.age = 21
	s6.grade = 2

	fmt.Println("name  :", s6.name)
    fmt.Println("age   :", s6.age) // s6.person.age
    fmt.Println("grade :", s6.grade)

	//! Implementasi Embeded Struct Dengan Nama Property Yang Sama
	var s7 = scientist{}
    s7.name = "wick"
    s7.age = 21        // age of student
    s7.person.age = 22 // age of person

    fmt.Println(s7.name)
    fmt.Println(s7.age)
    fmt.Println(s7.person.age)

	//! Pengisian Nilai Sub-Struct
	var p1 = person{name: "wick", age: 21}
	var s8 = scientist{person: p1, age: 10, grade: 2}

	fmt.Println("name  :", s8.name)
	fmt.Println("age   :", s8.age)
	fmt.Println("grade :", s8.grade)

	//! Anonymous Struct
	var s9 = struct {
		person
		grade int
	}{}
	
	s9.person = person{"wick", 21}
	s9.grade = 2

	fmt.Println("name  :", s9.person.name)
	fmt.Println("age   :", s9.person.age)
	fmt.Println("grade :", s9.grade)

	// anonymous struct tanpa pengisian property
	var _ = struct {
		person
		grade int
	}{}

	// anonymous struct dengan pengisian property
	var _ = struct {
		person
		grade int
	}{
		person: person{"wick", 20},
		grade: 2,
	}

	//! Kombinasi Slice & Struct
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 23},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	//! Inisialisasi Slice Anonymous Struct
	var students = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"wick", 21}, grade: 2},
		{person: person{"wick", 21}, grade: 2},
	}

	for _, student := range students {
		fmt.Println(student)
	}

	//! Deklarasi Anonymous Struct Menggunakan Keyword var
	// dicetak sebuah objek dari anonymous struct 
	// kemudian disimpan pada variabel bernama student
	var _ struct {
		grade int
	}

	// deklarasi sekaligus inisialisasi
	var _ = struct {
		grade int
	}{
		12,
	}

	//! Nested Struct
	type _ struct {
		person struct {
			name string
			age	 int
		}
		grade int
		hobbies []string
	}

	//! Deklarasi & Inisialisasi Struct Secara Horizontal
	type _ struct { name string; age int; hobbies []string }

	var _ = struct { name string; age int; } { age: 22, name: "wick" }
	var _ = struct { name string; age int; } { "ethan", 23 }

	//! Tag Property Dalam Struct
	type _ struct {
		name string `tag1`
		age  int	`tag2`
	}

	//! Type Alias
	type Person struct {
		name string
		age  int
	}
	type People = Person

	var v1 =  Person{"wick", 21}
	fmt.Println(v1)

	var v2 =  Person{"john", 21}
	fmt.Println(v2)

	people := People{"wick", 21}
	fmt.Println(Person(people))

	person := Person{"wick", 21}
	fmt.Println(People(person))

	type _ = struct {
		name string
		age  int
	}

	type Number = int
	var _ Number =  12
}

//! Deklarasi Struct
type student struct {
	name  string
	grade int
}

//! Embeded Struct
type person struct {
	name string
	age  int
}

type scientist struct {
	age int
	grade int
	person
}