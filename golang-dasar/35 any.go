package main

import "fmt"

//	interface kosong atau any adalah interface yang tidak memiliki deklarasi method satupun,
//	hal ini membuat secara otomatis semua tipe data akan menjadi implemantasinya
//	beberapa contoh penggunaan interface : fmt.Println(a...interface{}), panic(v interface{}), recover() interface{}

func Ups() any {
	return "Ups"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
