package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println("Nilai int32 =", nilai32)
	fmt.Println("Nilai int64 =", nilai64)
	fmt.Println("Nilai int16 =", nilai16)

	var name = "Hilmi Yahya"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println("Name =", name)
	fmt.Println("e =", e)
	fmt.Println("eString =", eString)
}