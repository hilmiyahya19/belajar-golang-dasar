package main

import "fmt"

func main() {
	name := "Hilmi"

	switch name {
	case "Hilmi":
		fmt.Println("Hello Hilmi")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hello, who are you?")
	}

	name = "Yahya"
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	name = "HilmiYahyaa"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}