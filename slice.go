package main

import "fmt"

func main() {
	names := [...]string{"Eko", "Kurniawan", "Khannedy", "Joko", "Budi", "Nugraha"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	var slice3 []string = names[3:]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1) // [Sabtu, Minggu]

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1) // [Sabtu Baru, Minggu Baru]
	fmt.Println(days) // [Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"}
	fmt.Println(daySlice1) // [Sabtu Baru, Minggu Baru]
	fmt.Println(daySlice2) // [Sabtu Baru, Minggu Baru, Libur Baru]
	fmt.Println(days) // [Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru]

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Hilmi"
	newSlice[1] = "Hilmi"
	// newSlice[2] = "Hilmi" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Hilmi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Yahya"
	fmt.Println(newSlice2)
	fmt.Println(newSlice) // berubah (ada elemen "Yahya"), karena masih menggunakan array yg sama

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5} // Array menyertakan length sebagai bagian dari tipe
	iniArray2 := [5]int{1, 2, 3, 4, 5} // Array menyertakan length sebagai bagian dari tipe
	iniSlice := []int{1, 2, 3, 4, 5} // Slice memiliki length tetapi tidak menjadi bagian dari tipe

	fmt.Println(iniArray)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)
}	