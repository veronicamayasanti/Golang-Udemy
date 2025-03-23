package main

import "fmt"

//	Operator new
//	selain menggunakan operator & untuk membuat pointer
//	bisa juga menggunakan function new,namun hanya mengembalikan data kosong,tidak ada data awal

type Address struct {
	City, Province, Country string
}

func main() {
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "New York"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
