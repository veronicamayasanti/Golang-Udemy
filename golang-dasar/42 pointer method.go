package main

//	pointer method
//	walaupun method akan menempel di struct,tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
//	sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memnaggil method

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
