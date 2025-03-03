package main

//	bisa membuat variable secara langsung di tipe data return functionnya
import "fmt"

func getCompleteName() (firstName, middleName, lastName string, age int) {
	firstName = "veronica"
	middleName = "maya"
	lastName = "santi"
	age = 33

	return firstName, middleName, lastName, age
}

func main() {
	firstName, middleName, lastName, age := getCompleteName()
	fmt.Println(firstName, middleName, lastName, age)
}
