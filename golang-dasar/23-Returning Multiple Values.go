package main

import "fmt"

//	returning multiple values
//	function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value

func getFullName() (string, string) { //	untuk memberitahu jika function mengembalikan multiple value,
	// harus menulis semua tipe data return valuenya di function

	return "maya", "santi"
}

func main() {
	//first, second := getFullName()
	//	multiple return value wajib ditangkap semua valuenya
	//	jika ingin menghiraukan return value tersebut, bisa menggunakan tanda _ (garis bawah)
	first, _ := getFullName()
	fmt.Println(first)
}
