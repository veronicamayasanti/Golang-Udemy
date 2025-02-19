package main

import "fmt"

func main() {
	//	tipe data array
	//	array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
	//	saat membuat array, perlu menentukan jumlah data yang bisa ditampung oleh array tersebut
	//	daya tampung array tidak bisa bertambah setelah array dibuat

	// index di array
	// 0, 1, 2 = 1, 2, 3

	var name [3]string
	name[0] = "A"
	name[1] = "B"
	name[2] = "C"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	//	cara membuat array secara langsung saat deklarasi variable
	var values = [3]int{
		90, 70, 50,
	}

	fmt.Println("deklarasi array secara langsung", values)
	fmt.Println("nilai values index ke 0 ", values[0])
	fmt.Println("nilai values index ke 1 ", values[1])
	fmt.Println("nilai values index ke 2 ", values[2])

	//	kode program array
	//	menggunakan ... titik 3 jika tidak ingin menentukan jumlah arraynya namun harus di deklarasikan terlebih dahulu
	//	operasi			keterangan
	//	len(array)		untuk mendapatkan panjang array
	//	array[index]	mendapat data di posisi index
	//	array[index] = value	mengubah data di posisi index

	var numbers = [...]int{
		1,
		2,
		3,
		4,
	}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	numbers[1] = 100
	fmt.Println(numbers)

}
