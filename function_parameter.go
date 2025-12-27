package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Hilmi", "Yahya")
	sayHelloTo("Hilmi", "Yahya")
}