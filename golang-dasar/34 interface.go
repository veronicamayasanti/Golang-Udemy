package main

import "fmt"

//	interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
//	sebuah interface berisikan definisi-definisi method
//	interface digunakan sebagai kontrak

//	implementasi interface
//	setiap data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
//	sehingga tidak perlu mengimplementasikan interface secara manual

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "John"}
	SayHello(person)

	animal := Animal{"kucing"}
	SayHello(animal)
}
