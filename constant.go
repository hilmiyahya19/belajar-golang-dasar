package main

import "fmt"

func main() {
	const (
		firstName string = "Hilmi"
		lastName = "Yahya"
	)

	// error
	// firstName = "Nama Baru"
	// lastName = "Nama Baru"

	fmt.Println(firstName)
	fmt.Println(lastName)
}