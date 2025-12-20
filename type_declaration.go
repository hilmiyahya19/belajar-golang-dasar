package main

import "fmt"

func main() {
	
	type NoKTP string

	var ktpHilmi NoKTP = "11111111111111"

	var contoh string = "222222222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpHilmi)
	fmt.Println(contohKtp)
}