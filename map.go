package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Hilmi"
	// person["address"] = "Kediri"

	person := map[string]string{
		"name": "Hilmi", 
		"address": "Kediri",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Hilmi"
	book["ups"] = "salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}