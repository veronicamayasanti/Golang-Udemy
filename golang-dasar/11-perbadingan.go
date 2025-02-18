package main

import "fmt"

func main() {
	//	operasi Perbandingan
	//	adalah operasi untuk membandingkan dua buah data
	//	adalah operasi yang menghasilkan nilai boolean (benar atau salah)
	//	jika hasil operasonya adalah benar maka nilainya adalah true, jika hasil sebaliknya maka nilainya false
	//	operator	keterangan
	//	>			lebih dari
	//	<			kurang dari
	//	>=			lebih dari sama dengan
	//	<=			kurang dari sama dengan
	//	==			sama dengan
	//	!=			tidak sama dengan

	name1 := "maya"
	name2 := "maya"
	result := name1 == name2

	angka1 := 1
	angka2 := 11
	hasilnya := angka1 == angka2
	fmt.Println("hasil perbandingan name1 == name2 adalah", result)
	fmt.Println(hasilnya)

}
