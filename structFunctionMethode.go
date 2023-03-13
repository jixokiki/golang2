package main

import "fmt"

type Customer struct {
	Nama, Alamat string
	Age, Number  int
}

//tanpa menggunakan struct methode function
func sayHi(pembeli Customer, name string) {
	fmt.Println("Hello", name, "My Name Is", pembeli.Nama)
}

//menggunakan struct methode function
func (penjual Customer) sayH(name string) {
	fmt.Println("Hello", name, "My Name Is", penjual.Nama)
}

func main() {
	pembeli := Customer{
		Nama: "Rizki Maulana",
	}
	sayHi(pembeli, "Hito")

	pembeli.sayH("Owen")
}
