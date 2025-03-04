package main

import "fmt"

//	variadic function
//	parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
//	varargs artinya bisa menerima lebih dari satu input, semacam Array
//	perbedaan dengan parameter biasa dengan tipe data Array
//	jika parameter tipe Array, wajib membuat Array terlebih dahulu sebelum mengirimkan ke function
//	jika parameter menggunakan varargs, bisa langsung mengirim datanya, jika lebih dari satu cukup menggunakan tanda koma

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//	slice parameter
//	menggunakan variadic function namun memiliki variable berupa slice
//	bisa menjadikan slice sebagai vararg parameter

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println(total)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
