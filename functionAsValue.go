package main

import "fmt"

func getName(n string) string {
	return "Wellcome " + n
}

func main() {
	nama := getName

	fmt.Println(nama("Iki"))

	//klo mau dimasukan dulu ke parameter
	gabungan := nama("Rizki")
	fmt.Println(gabungan)
}
