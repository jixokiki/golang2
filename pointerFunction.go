package main

import "fmt"

type Address struct {
	City, Provincy, Country string
}

func ChangeAddress(tujuan *Address) {
	tujuan.City = "Bogor"
	tujuan.Country = "Jawa Timur"

}

func main() {
	alamat := Address{
		City:     "Bekasi",
		Provincy: "Jawa Barat",
		Country:  "Indonesia",
	}
	fmt.Println(alamat)

	ChangeAddress(&alamat) //ini jika belum mempunyai variabel untuk alamat untuk menyatakan pointer
	fmt.Println(alamat)

	//klo sudah ada variabel alamatnya
	var alamatPointer *Address = &alamat
	fmt.Println(alamatPointer)
}
