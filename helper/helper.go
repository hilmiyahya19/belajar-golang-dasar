package helper

import "fmt"

var version = "1.0.0" // Private variable (ga bisa diakses diluar package)
var Application = "Golang"

// Private function (ga bisa diakses diluar package)
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Hilmi")
	fmt.Println(version)
}