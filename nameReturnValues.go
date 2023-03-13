package main

import "fmt"

func hello() (first, last string) {
	first = "iki"
	last = "maulana"
	return
}
func main() {

	a, b := hello()
	fmt.Println(a, b)
}
