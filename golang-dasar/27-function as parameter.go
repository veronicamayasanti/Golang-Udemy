package main

import "fmt"

// function as parameter
// function bisa digunakan sebagai parameter untuk function lain

//func sayHelloWithFilter(name string, filter func(string) string) {
//	fmt.Println("Hello " + filter(name))
//}

// type declarations juga bisa digunakan untuk membuat alias function
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello " + filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("anjing", spamFilter)
	sayHelloWithFilter("maya", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("santi", filter)
	sayHelloWithFilter("anjing", filter)
}
