package main

import "fmt"

func main() {
	name := "Hilmi"
	fmt.Println(name)

	name = "Yahya"
	fmt.Println(name)

	var (
		firstName = "Hilmi"
		lastName = "Yahya"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}