package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("Hello/World.go"))
	fmt.Println(path.Base("Hello/World.go"))
	fmt.Println(path.Ext("Hello/World.go"))
	fmt.Println(path.Join("Hello", "world", "main.go"))
}
