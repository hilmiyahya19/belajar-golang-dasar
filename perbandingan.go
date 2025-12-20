package main

import "fmt"

func main() {
	var name1 = "Hilmi"
	var name2 = "Hilmi"

	var result bool = name1 == name2
	fmt.Println(result)

	var huruf1 = "b" // b > a // anggap aja b = 2, a = 1 sesuai abjad
	var huruf2 = "a"

	var result2 = huruf1 > huruf2
	fmt.Println(result2)
}