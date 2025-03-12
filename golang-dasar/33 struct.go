package main

import "fmt"

//	struct adalah sebuah template data yang digunakan untuk menggabungkan no atau lebih tipe data lainnya dalam satu kesatuan
//	struct biasanya representasi data dalam program aplikasi yang kita buat
//	data di struct disimpan dalam field
//	struct adalah kumpulan dari field

type Customer struct {
	Name, Address string
	Age           int
}

//	struct adalah template data atau prototype data
//	struct tidak bisa langsung digunakan
//	namun kita bisa membuat data/object dari struct yang telah dibuat

func (customer Customer) sayHello(name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func main() {
	var maya Customer
	maya.Name = "maya dev"
	maya.Address = "tendjo"
	maya.Age = 33

	fmt.Println(maya)
	fmt.Println(maya.Name)
	fmt.Println(maya.Address)
	fmt.Println(maya.Age)

	santi := Customer{
		Name:    "santi",
		Address: "bogor",
		Age:     22,
	}
	fmt.Println(santi)

	veronica := Customer{"veronica", "jakarta", 34}
	fmt.Println(veronica)

	santi.sayHello("samantha")
}
