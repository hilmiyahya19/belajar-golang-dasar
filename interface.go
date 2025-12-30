package main

import "fmt"

// Interface HasName
// Berisi kontrak: semua tipe yang ingin dianggap "HasName"
// HARUS memiliki method getName() yang mengembalikan string
type HasName interface {
	getName() string
}

// Function sayHello menerima parameter bertipe interface HasName
// Artinya: fungsi ini bisa menerima nilai dari tipe APAPUN
// selama tipe tersebut memiliki method getName()
func sayHello(value HasName) {
	fmt.Println("Hello", value.getName())
}

// Struct Person memiliki field Name
type Person struct {
	Name string
}

// Method getName milik Person
// Karena Person memiliki method getName() string,
// maka Person SECARA OTOMATIS mengimplementasikan interface HasName
func (person Person) getName() string {
	return person.Name
}

// Struct Animal memiliki field Name
type Animal struct {
	Name string
}

// Method getName milik Animal
// Karena Animal juga memiliki method getName() string,
// maka Animal juga SECARA OTOMATIS mengimplementasikan interface HasName
func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	// Membuat object Person
	person := Person{Name: "Hilmi"}

	// Person bisa dikirim ke sayHello
	// karena Person mengimplementasikan interface HasName
	sayHello(person)

	// Membuat object Animal
	animal := Animal{Name: "Kelinci"}

	// Animal juga bisa dikirim ke sayHello
	// karena Animal juga mengimplementasikan interface HasName
	sayHello(animal)
}
a