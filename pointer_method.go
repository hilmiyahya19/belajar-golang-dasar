package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
}

func main() {
	Eko := Man{"Eko"}
	Eko.Married()
	fmt.Println(Eko.name)
}