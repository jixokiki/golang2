package main

import "fmt"
func main(){
	var months = [...] string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Saturnus"
	fmt.Println(slice1)

	slice1[0] = "Uranus"
	fmt.Println(months)

	//var slice2 = months[2:4] (Jika array nya itu mencukupi menampung data maka akan dimasukan dan dapat mengubah array aslinya, tetapi jika tidak maka akan membuat array baru dan tidak akan menggangu array yang lama)
	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Jupiter")
	fmt.Println(slice3)

	slice3[1] = "Neptunus"
	fmt.Println(slice3)

	fmt.Println(months)
	fmt.Println(slice1)
	fmt.Println(slice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rizki"
	newSlice[1] = "Maulana"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice1 := make([]string, 2,5)
	newSlice1[0] = "Hito"
	newSlice1[1] = "Jeanet"

	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	copySlice := make([]string, len(newSlice1), cap(newSlice1))
	copy(copySlice, newSlice1)
	fmt.Println(copySlice)

	copySlice1 := make([]string, len(newSlice1), cap(newSlice))
	copy(copySlice1, newSlice)
	fmt.Println(copySlice1)

	copySlice2 := make([]string, 1, cap(newSlice1))
	copy(copySlice2, newSlice1)
	fmt.Println(copySlice2)
}