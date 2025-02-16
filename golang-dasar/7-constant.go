package main

import "fmt"

func main() {
	//	constant
	//	adalah variable yang nilainya tidak bisa diubah setelah pertama kali diberi nilai
	//	menggunakan kata kunci const
	//	wajib langsung menginisialisasikan datanya

	const firstName string = "veronica"
	const lastName = "maya"

	//	 akan error jika nilainya diubah
	// lastName = "wijaya"

	fmt.Println(firstName)
	fmt.Println(lastName)

	//	cara deklarasi multiple constant
	const (
		age     = 33
		address = "bogor"
	)
	fmt.Println(age)
	fmt.Println(address)
}
