package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

func main() {
	budi := Man{"Budi"}
	budi.Married()

	fmt.Println(budi.Name)
}
