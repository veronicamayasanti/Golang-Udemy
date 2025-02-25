package main

import "fmt"

func main() {
	// if expression
	// if adalah salah satu kata kunci yang digunakan untuk percabangan
	// percabangan artinya bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi

	//	else expression
	// blok if akan dieksekusi ketika kondisi if bernilai true
	//	untuk melakukan eksekusi program tertentu jika kondisi if bernilai false bisa menggunakan else expression

	//	else if expression
	//	untuk membuat beberaa kondisi menggunakan else if exression

	name := "veronica"
	if name == "maya" {
		fmt.Println("hello maya")
	} else if name == "santi" {
		fmt.Println("hello santi")
	} else if name == "veronica" {
		fmt.Println("hello veronica")
	} else {
		fmt.Println("hello siapa nama mu")
	}

	//	if dengan short statement
	// if mendukung short statement sebelum kondisi

	if length := len(name); length > 3 {
		fmt.Println("name is too long")
	} else {
		fmt.Println("name is too short")
	}
}
