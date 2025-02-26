package main

import "fmt"

func main() {
	//	for loops adalah salah satu fitur perulangan

	//counter := 1
	//for counter <= 10 {
	//	fmt.Println("perulangan ke - ", counter)
	//	counter++
	//}
	//fmt.Println("selesai")

	//	for dengan statement
	//	dalam for bisa menambahkan statement, terdapat 2 statement yang bisa tambahkan di for
	//	init statement yaitu statement sebelum for dieksekusi
	//	post statement yaitu statemen yang akan selalu dieksekusi di akhir tiap perulangan

	for counter := 1; counter <= 8; counter++ {
		fmt.Println("perulangan ke - ", counter)

	}
	fmt.Println("sampai sini perulangannya")

	//	for range
	//	for bisa digunakan untuk melakukan iterasi terhadap semua data collection
	//		data collection contohnya Array, Slice, dan Map

	names := []string{"geo", "tata", "maya", "alex"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
