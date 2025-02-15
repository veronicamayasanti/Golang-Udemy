// tipe data string
// adalah tipe data kumpulan karakter
// jumlah karakter  di dalam string bisa nol sampai tidak terhingga
// direpresentasikan dengan kata kunci string
// selalu diawali dengan karakter " dan diakhiri "

package main

import "fmt"

func main() {
	fmt.Println("veronica")
	fmt.Println(" veronica maya")
	fmt.Println(" veronica maya santi")

	//	function untuk string
	// function				keterangan
	// len("string")		menghitung jumlah karakter di string
	// "string"[number]		mengambil karakter pada posisi yang di tentukan

	fmt.Println(len("veronica"))
	fmt.Println(" veronica maya"[0])
	fmt.Println(" veronica maya santi"[1])
}
