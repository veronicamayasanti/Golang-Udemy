package main

import "fmt"

//	data nil digunakan di beberapa tipe data seperti interface, function, map, slice, pointer dan channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("maya")

	if data == nil {
		fmt.Println("data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
