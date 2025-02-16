package main

import "fmt"

func main() {
	//	konversi tipe data
	//	bisa melakukan konversi tipe data dari satu tipe ke tipe lain
	//	misal dari tipe data int32 ke int63 dll
	//	contoh 1
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//	contoh 2
	var name = "veronica maya"
	var e = name[0]
	var eString = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
