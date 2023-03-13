package main

import "fmt"

type HisName interface {
	GetName() string
}

func spillName(hisnama HisName) {
	fmt.Println("Haii", hisnama.GetName())
}

type TheName struct {
	Name string
}

func (thename TheName) GetName() string {
	return thename.Name
}

type NIM interface {
	GetNim() int
}

func KetNim(nim NIM) {
	fmt.Println(nim.GetNim(), "Login")
}

type TheNim struct {
	NIM int
}

func (thenim TheNim) GetNim() int {
	return thenim.NIM
}

func main() {
	var iki TheName
	iki.Name = "Rizki"

	var nim_ TheNim
	nim_.NIM = 20200801041
	KetNim(nim_)
	spillName(iki)

}
