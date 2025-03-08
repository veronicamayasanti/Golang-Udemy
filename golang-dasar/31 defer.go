package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication(error bool) {
	//	defer function adalah function yang bisa dijadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
	//	akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

	defer logging()
	//fmt.Println("Run Application")

	//	panic function adalah function yang bisa digunakan untuk menghentikan program
	//	saat panic function dipanggil, program akan berhenti, namun defer function tetap akan dieksekusi

	if error {
		panic("error")
	}

	//	recover
	//	recover adalah function yang digunakan untuk menangkap data panic
	//	dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

	//cara salah
	//message := recover()
	//fmt.Println("terjadi error = ", message)

}

func endApp() {
	fmt.Println("endApp")
	message := recover()
	fmt.Println("terjadi error", message)

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("error")
	}
}

func main() {
	runApplication(false)

	runApp(false)

}
