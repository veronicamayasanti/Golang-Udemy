package helper

import "fmt"

var version = "1.0.0" // tidak bisa di akses
var Application = "golang"

// tidak bisa di akses di luar package karena huruf depan kecil
func sayGoodBye(nama string) string {
	return "Hello " + nama
}

// bisa di akses karena huruf depan function besaqr
func SayHello(name string) string {
	return "Hello " + name
}

// sayGoodBye bisa di akses dalam package yang sama
func Contoh() {
	sayGoodBye("santi")
	fmt.Println(version)
}
