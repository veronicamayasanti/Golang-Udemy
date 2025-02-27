package main

import "fmt"

//	function bisa mengembalikan data

func getHello(name string) string { // untuk memberi tahu bahwa function mengembalikan data, harus menulikan tipe data kembalian dari function tersebut
	//	untuk mengembalikan data dari function menggunakan kata kunci return di ikuti dengan datanya
	return "Hello " + name
}

func sum(num int) int {
	return num + num
}

func main() {
	//result := getHello("Bob")
	//fmt.Println(result)

	resSum := sum(11)
	fmt.Println(resSum)
}
