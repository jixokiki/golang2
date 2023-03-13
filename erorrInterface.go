package main

import (
	"errors"
	"fmt"
)

// kita akan melakukan pengecekan eror pada function pembagian
func bagi(nilai1 int, nilaiPembagi int) (int, error) {
	if nilaiPembagi == 0 {
		return 0, errors.New("Pembagi Tidak Boleh 0")
	} else {
		result := nilai1 / nilaiPembagi
		return result, nil
	}
}

/*
*
sebenernya di golang sudah disediakan package sendiri buat error yaitu sebagai berikut
*/
func main() {
	var eror error = errors.New("Ups You Have Trouble")
	fmt.Println(eror)

	hasil, err := bagi(100, 10)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
}
