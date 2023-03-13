package main
import "fmt"

func main(){
	for i:= 0; i <10; i++{
		if i == 5{
			break
		}
		fmt.Println("Urutan :", i)
	}

	for a:= 0; a <10; a++{
		if a == 7{
			continue
		}
		fmt.Println("Urutan :", a)
	}
}