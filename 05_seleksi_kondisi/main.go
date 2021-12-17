package main

import "fmt"

func main() {
	//! Seleksi Kondisi Dengan Menggunkan Keyword if, else if, & else
	var point = 10

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	//! Variabel Temporary if - else
	var value = 8840.0

	if percent := value / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

	//! Seleksi Kondisi Menggunakan Keyword switch - case
	var balance = 6

	switch balance {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//! Pemanfaatan case Untuk Banyak Kondisi
	var number = 5

	switch number {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	//! Switch dengan daya if - else
	var power = 6

	switch {
	case power == 8:
		fmt.Println("perfect")
	case (power < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	//! Penggunaan Keyword fallthrough Dalam switch
	var speed = 5

	switch {
	case speed == 8:
		fmt.Println("perfect")
	case (speed < 8) && (speed > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	//! Seleksi Kondisi Bersarang
	var health = 2

	if health > 7 {
		switch health {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if health == 5 {
			fmt.Println("not bad")
		} else if health == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if health == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}