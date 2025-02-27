package main

//	faction parameter
//	bisa menmabhka paremater/data dari luar ke dala function
//	parameter tidak wajib
//	ketika membuat function dengan paramter maka ketika memanggil function wajib memasukan datanya

import "fmt"

func sayHelloo(firstName string, lastName string) {
	fmt.Println("Hello " + firstName + " " + lastName)
}

func main() {
	sayHelloo("maya", "santi")
}
