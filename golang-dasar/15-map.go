package main

import "fmt"

func main() {
	//	tipe data map
	//	map adalah tipe data lain yang berisikan kumpulan data yang sama
	//	bisa menentukan jenis tipe data index yang akan digunakan
	//	map adalah tipe data kumpulan key-value (kata kunci - nilai)
	//	kata kunci bersifat unik dan tidak boleh sama dengan kata kunci yang lain

	person := map[string]string{
		"name": "maya",
		"age":  "33",
	}
	fmt.Println("this is data person : ", person)
	fmt.Println("this is data person name : ", person["name"])
	fmt.Println("this is data person age : ", person["age"])
	fmt.Println("this is data lenght person : ", len(person))

	//	function map
	//	len(map) = untuk mendapatkan jumlah data di map
	//	map[key] = mengambil data di map dengan key
	//	map[key] = value	: mengubah data di map dengan key
	//	make(map[TypeKey]TypeValue)	= membuat map baru
	//	delete(map,key)

	book := make(map[string]string)
	book["title"] = "My Book"
	book["description"] = "This is an example of book"
	book["author"] = "My Author"
	book["wrong"] = "ups"
	delete(book, "wrong")
	fmt.Println(book)

}
