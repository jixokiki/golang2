package main

import "fmt"

func kelas() interface{} {
	return 0
}

func main() {
	//cara pertama pengubahan type asertion
	//var pengubah interface{} = kelas()
	//var pengubahString string = pengubah.(string)

	//fmt.Println(pengubahString)

	//cara kedua pengubahan type asertion menggunakan switch (cara aman)

	var result interface{} = kelas()
	switch hasil := result.(type) {
	case int:
		fmt.Println("Hasil", hasil, "is int")
	case string:
		fmt.Println("Hasil", hasil, "is string")
	case bool:
		fmt.Println("Hasil", hasil, "is boolean")
	default:
		fmt.Println("Uknown Type")
	}
}
