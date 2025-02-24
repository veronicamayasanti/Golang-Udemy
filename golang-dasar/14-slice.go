package main

import "fmt"

func main() {
	//	tipe data slice
	//	tipe data slice adalah potongan dari data Array
	//	ukuran slice bisa berubah
	//	slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di Array

	//	tipe data slice memiliki 3 data:
	//	pointer adalah penunjuk data pertama di array para slice
	//	length adalah panjang dari slice
	//	capacity adalah kapasitas dari slice, length tidak boleh lebih banyak dari capacity

	//	membuat slice dari array
	//	membuat slice			keterangan
	//	array[low:high]			membuat slice dari array dimulai index low sampai index sebelum high
	//	array[low:]				membuat index dari array dimulai index low sampai index akhir di array.
	//	array[:high]			membuat slice dari array dimulai index 0 sampai index sebelum high
	//	array[:]				membuat slice dari array dimulai index 0 sampai index akhir di array

	//	slice dan array.
	//	array { 0=jan, 1=feb, 2=mar, 3=apr, 4=mei, 5=jun, 6=jul, 7=agus, 8=sep, 9=okto, 10=nov, 11=des]
	//	array[4:7] slice : pointer =4, length=3, capacity = 8.
	//	array[6:9] slice: pointer=6, length=3, capacity= 6.

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	slice1 := days[2:5]
	slice2 := days[5:]
	slice3 := days[4:]
	slice4 := days[:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	//	function slice
	//	operasi			keterangan
	//	len(slice)		untuk mendapatkan panjang
	//	cap(slice)		untuk mendapat kapasitas
	//	append(slice, data)		membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru
	//	make([]TypeData, length, capacity)		membuat slice baru
	//	copy(destination, source)	menyalin slice dari source ke destination

	numbers := [...]string{"one", "two", "three", "four", "five", "six", "seven"}
	// array[low:]
	//Membuat slide dari array dimulai index low sampai index akhir di array
	numbers5 := numbers[4:]
	numbers5[0] = "five-update"
	fmt.Println("this ini slice numbers = ", numbers)
	fmt.Println("this ini slice [4:] = ", numbers5)
	fmt.Println("this ini slice [0] update =", numbers5)

	numbersNew := append(numbers5, "delapan", "sembilan", "sepuluh")
	fmt.Println("this ini slice numbersNew = ", numbersNew)
	numbersNew[0] = "lima"
	fmt.Println("this ini slice numbersNew [0] update = ", numbersNew)
	fmt.Println(numbers)

	//	make([]TypeData, length, capacity) : Membuat slice baru
	newSlice := make([]string, 3, 5) // 3= panjang slice, 5=kapasitas slice
	newSlice[0] = "one"
	newSlice[1] = "two"
	newSlice[2] = "three"
	fmt.Println("make newslice : ", newSlice)
	fmt.Println("length newSlice = ", len(newSlice))
	fmt.Println("capacity newSlice = ", cap(newSlice))

	//	 copy(destination, source)
	//Menyalin slice dari source ke destination
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("fromSlice : ", fromSlice)
	fmt.Println("toSlice : ", toSlice)

	//membedakan array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("iniArray : ", iniArray)
	fmt.Println("iniSlice : ", iniSlice)
}
