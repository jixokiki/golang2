package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func main() {
	tempat := Alamat{"Bekasi", "Jawa Barat", "Indonesia"}
	tempat2 := tempat
	tempat3 := &tempat // supaya menjadi pointer hanya menambahkan & dan ini merubah data pada data yang sudah ada

	tempat4 := *&tempat
	tempat2.City = "Bandung"
	tempat3.City = "Surabaya"
	tempat4.City = "Yogyakarta"

	var tempat5 *Alamat = &Alamat{"Depok", "Jawa Barat", "Indonesia"}
	*tempat3 = Alamat{"Bogor", "Jawa Barat", "Indonesia"}

	tempat6 := new(Alamat)
	tempat6.City = "Bekasi"
	fmt.Println(tempat)
	fmt.Println(tempat2)
	fmt.Println(tempat3)
	fmt.Println(tempat4)
	fmt.Println(tempat5)
	fmt.Println(tempat6)

}
