package main
import "fmt"
func main(){
	Nama1 := "Iki"
	Nama2 := "Rizki"

	result1 := Nama1 == Nama2
	result2 := Nama1 >= Nama2
	result3 := Nama1 <= Nama2
	result4 := Nama1 != Nama2
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	number1 := 100
	number2 := 101

	fmt.Println(number1 == number2)
	fmt.Println(number1 != number2)
	fmt.Println(number1 >= number2)
	fmt.Println(number1 <= number2)
	fmt.Println(number1 > number2)
	fmt.Println(number1 < number2)
}