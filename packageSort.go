package main

import (
	"fmt"
	"sort"
)

type UserInterface struct {
	Nama string
	Usia int
}

type getUserInterface []UserInterface

func (value getUserInterface) Len() int {
	return len(value)
}

func (value getUserInterface) Less(i, j int) bool {
	return value[i].Usia < value[j].Usia
}

func (value getUserInterface) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	user := []UserInterface{
		{"Rizki", 20},
		{"Maulana", 24},
		{"Owen", 400},
		{"Hito", 33},
	}

	fmt.Println(user)

	sort.Sort(getUserInterface(user))
	fmt.Println(user)
}
