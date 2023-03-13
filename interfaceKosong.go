package main

import "fmt"

func trial(i int) interface{} {
	if i == 1 {
		return "Rizki Maulana"
	} else if i == 2 {
		return 20200801041
	} else if i == 3 {
		return true
	} else {
		return "telah absen"
	}

}

func main() {
	var d interface{} = trial(3)
	n := trial(1)
	fmt.Println(n)
	fmt.Println(d)
}
