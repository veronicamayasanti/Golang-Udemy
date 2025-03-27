package main

//	Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error,
//	nama interface itu adalah error, yang terdapat di package errors

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi is zero")
	} else {
		return nilai / pembagi, nil
	}
}
func main() {
	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("Error ", err)
	}
}
