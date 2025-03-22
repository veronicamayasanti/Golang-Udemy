package main

import "fmt"

//	pass by value
//	secara default di Go-lang semua variableitu di passing by value, bukan reference
//	artinya: jika mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirm adalah duplikasivaluenyaf

//	pointer
//	pointer adalah kemampuan mendapat reference ke lokasi data data memory yang sama, tanpa menduplikat data yang sudah ada
//	dengan kemampuan pointer bisa membuat pass by reference

//	operator &
//	untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, menggunakan oerator & diikuti dengan nama variablenya

type Address struct {
	City, Province, Country string
}

func main() {
	//	pass by value
	//address1 := Address{"tenjo", "bogor", "indonesia"}
	//address2 := address1 // copy value
	//
	//address2.City = "tigaraksa"
	//fmt.Println(address1) // tidak berubah addres1
	//fmt.Println(address2) // city berubah jadi tigaraksa

	//	pointer
	address1 := Address{"tenjo", "bogor", "indonesia"}
	address2 := &address1 // pointer

	address2.City = "tigaraksa"
	fmt.Println(address1)
	fmt.Println(address2)
}
