package main

import "fmt"

//	recursive function adalah function yang memanggil function dirinya sendiri
//	factorial for loop

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	factorial := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println("hasil factorial manual ", factorial)
	fmt.Println("hasil dari factorialLoop =", (factorialLoop(10)))
	fmt.Println("hasil factorialRecursive =", (factorialRecursive(10)))
}
