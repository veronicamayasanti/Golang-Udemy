package main

import "fmt"

func main() {
	//	Type Declarations
	//	adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
	//	digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti

	type NoKTP string
	var ktpMaya NoKTP = "010101"
	fmt.Println(ktpMaya)
	fmt.Println(NoKTP("23232323"))

}
