package main

import "fmt"

func main() {
	//	Switch expression
	//	selain if expression untuk melakukan percabangan bisa menggunakan switch expression
	//	digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

	name := "veronica maya santi"
	switch name {
	case "maya":
		fmt.Println("Maya")
	case "santi":
		fmt.Println("Santi")
	case "veronica":
		fmt.Println("Veronica")
	default:
		fmt.Println("kenalan donk ")
	}

	//	Switch dengan short statement
	//	Switch juga mendukung short statement sebelum variable yang akan di cek kodisinya
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	//	Switch tanpa kondisi
	// kondisi di Switch expression tidak wajib
	//	jika tidak menggunakan kondisi di switch expression, bisa menambahkan kondisi tersebut di setiap casenya

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}

}
