package main

import "fmt"

// membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
// Hal tersebut dinamakan anonymous function, atau function tanpa nama

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "babi"
	}
	registerUser("babi", blacklist)
	registerUser("maya", blacklist)
	registerUser("babi", func(name string) bool {
		return name == "babi"
	})
}
