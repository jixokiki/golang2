package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Rizki")    // data1
	data.PushBack("Maulana")  // data2
	data.PushBack("Franklyn") // data3
	data.PushFront("Owen")    //ini akan menambah list data ke ata paling awal yaitu rizki

	//mengambil data depan
	fmt.Println(data.Front().Value)
	//mengambil data belakang
	fmt.Println(data.Back().Value)

	//mengambil data dari depan ke belakang
	for d := data.Front(); d != nil; d = d.Next() {
		fmt.Println(d.Value)
	}

	//mengambil data dari belakang ke depan
	for dd := data.Back(); dd != nil; dd = dd.Prev() {
		fmt.Println(dd.Value)
	}
}
