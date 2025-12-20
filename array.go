package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Hilmi"
	names[1] = "Yahya"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [...]int{
		70,
		80,
		90,
		100,
	}

	fmt.Println(values) // [70 80 90 100]
	fmt.Println(values[0]) // 70
	fmt.Println(values[1]) // 80
	fmt.Println(values[2]) // 90
	fmt.Println(values[3]) // 100

	fmt.Println(len(values)) // 4
	values[3] = 110
	fmt.Println(values) // [70 80 90 110]
}
