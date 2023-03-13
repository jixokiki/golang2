package main

import "fmt"

type Kampus struct {
	Nama string
	Age  int
}

func (campus *Kampus) Wellcome() {
	campus.Nama = "Wellcome to " + campus.Nama
	campus.Age = campus.Age
}

func main() {
	c := Kampus{"Esa Unggul University", 21}
	c.Wellcome()
	d := Kampus{"University Pelita Harapan", 21}
	d.Wellcome()
	fmt.Println(c.Nama)
	fmt.Println(d.Age)

}
