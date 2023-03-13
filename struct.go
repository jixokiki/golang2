package main

import "fmt"

type Data struct {
	Nama, Matkul string
	Age, NIM     int
}

func main() {
	var iki Data

	//cara pertama deklarasi atau mengisi data ke struct
	iki.Nama = "Rizki Maulana"
	iki.Matkul = "Game Developer"
	iki.NIM = 20200801041
	iki.Age = 20

	fmt.Println(iki)

	//cara kedua deklarasi atau mengisi data ke struct
	rizki := Data{
		Nama:   "Rizki Maulana",
		NIM:    20200801041,
		Age:    21,
		Matkul: "Game Developer",
	}

	fmt.Println(rizki)

	//cara ketiga deklarasi atau mengisi data ke struct (cara ini isi structnya harus sama dengan ketentuan struct yang sudah kamu buat diatas, jika tidak sama atau kurang dia akan eror)
	upil := Data{"Rizki Maulana", "Game Developer", 22, 20200801041}
	fmt.Println(upil)
}
