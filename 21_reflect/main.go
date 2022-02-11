package main

import (
	"fmt"
	"reflect"
)

func main() {
	//! Mencari Tipe Data & Value Menggunakan Reflect
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	//! Pengaksesan Nilai Dalam Bentuk interface{}
	fmt.Println("nilai variabel :", reflectValue.Interface())
	
	var _ = reflectValue.Interface().(int)

	// Ujicoba getPropertyInfo
	var s1 = &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()

	// Ujicoba SetName
	var s2 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue2 = reflect.ValueOf(s2)
	var method = reflectValue2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", s1.Name)
}

type student struct {
	Name  string
	Grade int
}

//! Pengaksesan Informasi Informasi Property Variabel Objek
func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	// Cek pointer atau tidak
	if reflectValue.Kind() == reflect.Ptr {
		// Ambil objek reflect aslinya
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
        fmt.Println("tipe data :", reflectType.Field(i).Type)
        fmt.Println("nilai     :", reflectValue.Field(i).Interface())
        fmt.Println("")
	}
}

//! Pengaksesan Informasi Method Variabel Objek
func (s *student) SetName(name string) {
	s.Name = name
}