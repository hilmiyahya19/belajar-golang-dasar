# Belajar Golang Dasar – Programmer Zaman Now

Repositori ini berisi rangkuman, catatan, dan kode praktik hasil belajar **Golang Dasar** dari channel YouTube **Programmer Zaman Now** oleh **Eko Kurniawan Khannedy**.
Seluruh materi diringkas dari *slide resmi Golang Dasar*  dan diimplementasikan ulang sebagai latihan di repository ini.

---

## 🎯 Tujuan Belajar

* Memahami fundamental bahasa Go dari nol.
* Menguasai dasar-dasar pemrograman Go seperti tipe data, variable, slice, map, struct, interface, pointers, hingga error handling.
* Membiasakan diri dengan workflow development Go (modules, build, run).
* Menyiapkan pondasi untuk belajar Go-Lang tingkat lanjut seperti Goroutine, Database, Web, dan Microservices.

---

## 🧩 Daftar Materi

### 📌 1. Pengenalan Golang

* Sejarah & alasan belajar Golang
* Kelebihan Go (sederhana, cepat, concurrency, GC, populer untuk backend & microservices)
* Cara install Go, cek versi, dan tools pendukung (VSCode, GoLand)

### 📌 2. Project & Module

* Cara membuat module:

  ```sh
  go mod init nama-module
  ```
* Struktur project dan cara menjalankan file Go
* Konsep single `main()` dalam satu module

### 📌 3. Program Dasar

* `package main`
* `func main()`
* Import `fmt` dan penggunaan `Println`

### 📌 4. Tipe Data

* Number (integer & float)
* Boolean
* String & function string
* Konversi tipe data
* Type declarations (alias)

### 📌 5. Variable & Constant

* Deklarasi var, tipe data, dan short declaration `:=`
* Multiple variable declaration
* Constant & multiple const

### 📌 6. Operasi Dasar

* Operasi matematika
* Augmented assignments
* Unary operator
* Operasi perbandingan
* Operasi boolean (AND, OR, NOT)

### 📌 7. Array, Slice, dan Map

* Pengenalan array (fixed size)
* Slice (pointer, length, capacity)
* Append, make, copy
* Map key-value, operasi dasar map

### 📌 8. Percabangan

* If, else, else if
* If short statement
* Switch (dengan & tanpa kondisi)

### 📌 9. Perulangan

* Perulangan `for`
* Init & post statement
* For-range
* Break & continue

### 📌 10. Function

* Function dasar
* Parameter & return value
* Multiple return value
* Named return value
* Variadic function
* Function value & function as parameter
* Function type
* Anonymous function
* Recursive function
* Closure

### 📌 11. Error Handling

* Defer
* Panic
* Recover
* Error interface & custom error

### 📌 12. Struct

* Membuat struct
* Struct literal
* Method pada struct

### 📌 13. Interface

* Konsep interface di Go
* Implementasi otomatis
* Interface kosong (`interface{}` / `any`)
* Type assertions & switch assertions

### 📌 14. Pointer

* Pass by value vs pass by reference
* Operator `&` dan `*`
* Pointer di function
* Pointer di method
* Keyword `new`

### 📌 15. Package & Import

* Membuat package
* Access modifier berdasarkan kapitalisasi
* Function `init`
* Blank identifier `_`

---

## 📁 Struktur Folder

Contoh struktur yang direkomendasikan:

```
/
├── src/
│   ├── 01-hello-world/
│   ├── 02-variable/
│   ├── 03-const/
│   ├── 04-type-data/
│   ├── 05-array-slice-map/
│   ├── 06-function/
│   ├── 07-struct/
│   ├── 08-interface/
│   ├── 09-pointer/
│   ├── 10-error-handling/
│   └── ...
├── README.md
└── go.mod
```

---

## 🚀 Cara Menjalankan

1. Install Go:
   [https://golang.org/](https://golang.org/)

2. Clone repository:

   ```sh
   git clone <url-repo-anda>
   cd <nama-folder>
   ```

3. Jalankan file Go:

   ```sh
   go run namafile.go
   ```

4. Atau build:

   ```sh
   go build
   ```

---

## 📚 Sumber Belajar

* Channel YouTube Programmer Zaman Now
* Slide resmi *Go-Lang Dasar* (file PPT) 
* [https://www.programmerzamannow.com](https://www.programmerzamannow.com)
* Dokumentasi resmi Go: [https://go.dev/doc/](https://go.dev/doc/)

---

## ✨ Catatan Tambahan

Repository ini sepenuhnya ditujukan sebagai dokumentasi belajar, latihan coding, dan persiapan sebelum masuk ke materi tingkat lanjut seperti:

* Goroutine & Concurrency
* Go Web & REST API
* Database SQL & NoSQL
* Clean Architecture
* Microservices with Go

---
