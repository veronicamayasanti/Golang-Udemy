package main

import "fmt"

//	membuat function yang bisa mengubah data asli parameter dan menjadikan parameter sebagai pointer menggunakan operator * di parameternya

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "japan"
}

func main() {
	address := Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
