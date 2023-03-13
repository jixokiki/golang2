package main

import (
	"fmt"
	"strconv"
)

func main() {
	//mengubah string menjadi boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("eror : ", err.Error())
	}

	//mengubah string menjadi Int
	integer, err := strconv.ParseInt("2000000000000", 16, 64)
	if err == nil {
		fmt.Println(integer)
	} else {
		fmt.Println("eror : ", err.Error())
	}

	//mengubah string dengan cara format boolean
	boolean1 := strconv.FormatBool(true)
	fmt.Println(boolean1)

	//mengubah string dengan cara format integer
	integer1 := strconv.FormatInt(20_000, 2)
	fmt.Println(integer1)

	//mengubah string dengan cara format float
	floatt := strconv.FormatFloat(300, 16, 10, 64)
	fmt.Println(floatt)

	//cara lama yang lebih simple itu bisa menggunakan Atoi / Itoa
	number, err := strconv.Atoi("10000000")
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	nomor := strconv.Itoa(10.0)
	fmt.Println(nomor)

}
