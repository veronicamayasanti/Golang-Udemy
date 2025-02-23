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

}
