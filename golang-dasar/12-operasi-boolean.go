package main

import "fmt"

func main() {
	//	operasi Boolean
	//	Operator	Keterangan
	//	&&			Dan
	//	||			Atau
	//	!			Kebalikan

	//	Operasi &&
	//	Nilai 1		Operator	Nilai 2		Hasil
	//	true		&&			true		true
	//	true		&&			false		false
	//	false		&&			true		false
	//	false		&&			false		false

	// Operasi ||
	//	Nilai1	Operator	Nilai2	Hasil
	//	true	||			true	true
	//	true	||			false	true
	//	false	||			true	true
	//	false	||			false	false

	//	Operasi !
	//	Operator	Nilai 2		Hasil
	//	!			true		false
	//	!			false		true
	nilaiAkhir := 90
	absensi := 80
	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi >= 80
	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

}
