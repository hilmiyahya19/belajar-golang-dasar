package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{
			Message: "ID Tidak Boleh Kosong",
		}
	}
	if id != "hilmi" {
		return &notFoundError{
			Message: "Data Tidak Ditemukan",
		}
	}
	// ok
	return nil
}
 
func main() {
	err := SaveData("hilmi", nil)
	if err != nil {
		// terjadi error (if else)
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		// terjadi error (switch case)
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	} else {
		// sukses
		fmt.Println("Data berhasil disimpan")
	}
}