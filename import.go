package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses diluar package
	// fmt.Println(helper.sayGoodBye("Hilmi")) // tidak bisa diakses diluar package
}