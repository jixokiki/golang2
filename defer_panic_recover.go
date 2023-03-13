package main

import (
	"fmt"
)

// penerapan defer function
func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApp(value int) {
	// ini tidak menggunakan defer function
	//logging() // tidak terpanggil
	// karena tidak menggunakan defer function apabila terjadi kesalahan pada function runApp nah data pada function logging tidak bisa diakses dan func runApp juga sama tidak bisa diakses akhirnya terjadi eror

	//ini menggunakan defer function
	defer logging() //jika menggunakan defer function function logging akan terpanggil apabila data dari func runApp benar ataupun salah
	fmt.Println("Run App")
	v := 10 / value
	fmt.Println("Result ", v)
}

func main() {
	//runApp(0)
	runApp(2)
	fmt.Println("\n")
	outApp(false)
	//fmt.Println("\n")
	//outApp(true)
	fmt.Println("\n")
	getLog(false)
	fmt.Println("\n")
	getLog(true)
}

// penerapan panic function
func logOut() {
	fmt.Println("Aplikasi Sedang Melakukan Proses")
}

func outApp(eror bool) {
	defer logOut()
	if eror {
		panic("ERROR!!!!")
	}
	fmt.Println("Aplikasi Telah Log Out")
}

// penerapan recover function
func log() {
	fmt.Println("Memuat Data")
	message := recover()
	fmt.Println("Password atau Username", message)
}

func getLog(login bool) {
	defer log()
	if login {
		panic("Tejadi Kesalahan")
	}
	fmt.Println("Loggin...")

}
