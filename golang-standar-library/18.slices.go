package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"John", "Paul", "George", "Ringo"}
	value := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(name))
	fmt.Println(slices.Max(name))
	fmt.Println(slices.Min(value))
	fmt.Println(slices.Max(value))
	fmt.Println(slices.Contains(name, "Paul"))
	fmt.Println(slices.Index(name, "George"))
	fmt.Println(slices.Index(name, "Ringo"))
}
