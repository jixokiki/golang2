package main

import "fmt"

func main() {

	nilai := 80
	anggota := "eko"
	counter := func() {
		anggota = "rizki"
		nilai++
	}

	counter()
	fmt.Println(anggota)
	fmt.Println(nilai)

}
