package main

import "fmt"

//	operator *
//	saat mengubah variable pointer, maka yang berubah hanya variable tersebut
//	semua variable yang mengacu ke data yang sama tidak akan berubah
//	jika ingin mengubah seluruh variable yang mengacu ke data tersebut, bisa menggunakan operator *

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"tenjo", "bogor", "indonesia"}
	address2 := &address1 // pointer

	address2.City = "tigaraksa"
	fmt.Println(address1)
	fmt.Println(address2)

	//address2 = &Address{"jakarta", "DKI JAKARTA", "indonesia"}
	*address2 = Address{"jakarta", "DKI JAKARTA", "indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
