package main

import "fmt"

func sumAll(num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 90, 80, 70))

	//jika kita sudah mempunyai parameter yang berbentuk slice kita juga bisa mendeklarasikannya kedalam variadic
	slice := []int{10, 10, 10, 10, 10}
	//masukan kedalam function sumAll dengan parameter totall
	totall := sumAll(slice...)
	fmt.Println(totall)
}
