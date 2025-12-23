package main

import "fmt"

func main() {
	name := "Eko"

	if name == "Hilmi" {
		fmt.Println("Hello Hilmi")
	} else if name == "Yahya" {
		fmt.Println("Hello Yahya")
	} else if name == "Eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hello, who are you?")
	}

	// length := len(name)
	// if length > 5 {
	// 	fmt.Println("Nama Terlalu Panjang")
	// } else {
	// 	fmt.Println("Nama Sudah Benar")
	// }

	// bisa dipersingkat / if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}