package main

import "fmt"

func main() {
	// variable
	// adalah tempat untuk menyimpan data
	// digunakan supaya bisa mengakses data yang sama dimanapun
	// di Go-Lang variable hanya bisa menyimpan tipe data yang sama
	// menggunakan kata kunci var diikuti nama variable lalu tipe datanya
	// contong: var name string

	var name string
	name = " veronica maya"
	fmt.Println(name)

	name = "santi wijaya"
	fmt.Println(name)

	//	membuat variable wajib menyebutkan tipe data variable
	// namun jika langsung menginisialisasikan data pada variablenya,
	// maka tidak wajib menyebutkan tipe data variablenya

	var address = "jakarta"
	fmt.Println(address)
	// kata kunci var saat membuat variable tidak wajib
	// asalkan saat membuat variable langsung menginisialisasi datanya
	// gunakan kata kunci := saat menginisialisasikan data pada variable

	username := "mayadev"
	fmt.Println(username)

	//	deklarasi multiple variable
	// bisa membuat variable secara banyak sekiligus

	var (
		firstName  = "veronica"
		middleName = "maya"
		lastName   = "santi"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
