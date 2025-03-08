package main

import "fmt"

//	closure adalah kemampuan sebuah function berinteraksi dengan data data disekitarnya dalam scope yang sama

func main() {
	counter := 0
	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
