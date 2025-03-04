package main

import "fmt"

//	function as value
//	function adalah first class citizen
//	function juga merupakan tipe data dan bisa disimpan di dalam variable

func getGoodBye(name string) string {
	return "good bye " + name
}

func main() {
	varGetGoodBye := getGoodBye
	fmt.Println(varGetGoodBye("maya"))
}
