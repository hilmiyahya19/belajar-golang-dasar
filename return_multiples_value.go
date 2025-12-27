package main

import "fmt"

func getFullName() (string, string) {
	return "Hilmi", "Yahya"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}