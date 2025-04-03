package main

//	package string berisikan function untuk memnipulasi tipe data string
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("maya santi", "ayya"))
	fmt.Println(strings.Split("maya santi", " "))
	fmt.Println(strings.ToLower("veronica maya SANTI"))
	fmt.Println(strings.ToUpper("veronica maya"))
	fmt.Println(strings.Trim("     maya      ", " "))
	fmt.Println(strings.ReplaceAll("tata geo", "geo", "wijaya"))
}
