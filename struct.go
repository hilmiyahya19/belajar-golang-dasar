package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var Hilmi Customer
	fmt.Println(Hilmi) // default value: { "", "", 0 }

	Hilmi.Name = "Hilmi"
	Hilmi.Address = "Kediri"
	Hilmi.Age = 22
	fmt.Println(Hilmi) // { "Hilmi", "Kediri", 22 }
	fmt.Println(Hilmi.Name)
	fmt.Println(Hilmi.Address)
	fmt.Println(Hilmi.Age)

	Joko := Customer{
		Name: "Joko",
		Address: "Jakarta",
		Age: 30,
	}
	fmt.Println(Joko)

	Budi := Customer{"Budi", "Bandung", 20}
	fmt.Println(Budi)

	Hilmi.sayHello("Eko")
	Budi.sayHello("Eko")
	Joko.sayHello("Eko")
}	

