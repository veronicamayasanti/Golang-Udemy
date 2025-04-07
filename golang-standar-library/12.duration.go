package main

import (
	"fmt"
	"time"
)

//	duration
//	saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
//	package time memiliki type Duration, yang sebenarnya adalah alias untuk int64
//	terdapat banyak method yang bisa digunakn untuk memanipulasi duration

func main() {
	var duration1 time.Duration = 100 * time.Millisecond
	var duration2 time.Duration = 500 * time.Second
	var duration3 time.Duration = duration2 - duration1

	fmt.Println(duration1, duration2)
	fmt.Println(duration3)
	fmt.Println("duration : %d\n ", duration3)
}
