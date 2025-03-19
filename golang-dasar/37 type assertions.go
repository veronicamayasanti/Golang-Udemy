package main

import "fmt"

//	type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
//	fitur ini digunakan ketika bertemu dengan data interface kosong
//	saat salah menggunakan type assertions, maka bisa berakibat terjadi panic
//	jika panic dan tidak ter-recover, maka otomatis program akan mati
//	agar lebih aman menggunakan switch expression untuk melakukan type assertions

func random() any {
	return "OK"
	//return 12f3
	//return false
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)
	//resultInt := result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)

	}
}
