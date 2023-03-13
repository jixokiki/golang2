package main

import "fmt"

func getFill(name string, filter func(string) string) {
	getName := filter(name)
	fmt.Println("", getName)
}

func toxicFill(name string) string {
	if name == "Kontol" {
		return "..."
	} else {
		return name
	}
}

func main() {
	//menggunakan function parameter
	getFill("Iki", toxicFill)
	getFill("Kontol", toxicFill)
	//tanpa menggunakan function parameter
	getN("Kontol")

	spam := getNTF("IKI")
	fmt.Println(spam)
}

//fillterin tanpa menggunakan cara function parameter

func getN(name string) {
	nameFill := name
	if name == "Kontol" {
		nameFill = "..."
	}
	fmt.Println(nameFill)
}

//function parameter menggunakan function type declaration

type Fill func(string) string

func getNT(NAMA string, filter Fill) {
	getfit := filter(NAMA)
	fmt.Println(getfit)
}

func getNTF(NAMA string) string {
	if NAMA == "MEMEK" {
		return "..."
	} else {
		return NAMA
	}
}
