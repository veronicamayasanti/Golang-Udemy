package main

import "fmt"

func main() {
	//	Break & Continue
	//	adalah kata kunci yang bisa digunakan dalam perulangan
	//	Break digunakan untuk menghentikan seluruh perulangan
	//	Continue digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("perulangan ke  ", i)
	}
}
