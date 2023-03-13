package main

import "fmt"

//loopingan tanpa recursive
func p(value int) int {
	hasil := 1
	for i := value; i > 0; i-- {
		hasil *= i
	}

	return hasil
}

func r(v int) int {
	if v == 1 {
		return 1
	} else {
		return v * r(v-1)
	}
}

func main() {
	loop := p(5)
	fmt.Println(loop)

	Recursive := r(5)
	fmt.Println(Recursive)
}
