package main
import "fmt"
func main(){
	name:= "Rizki"

	switch name{
	case "Rizki":
			fmt.Println("Hello Rizki")
	case "Hito":
			fmt.Println("Hello Hito")
	default:
			fmt.Println("Kenalan dong")
	}

	switch length:=len(name);length>5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	Panjang:= len(name)
	switch {
	case Panjang>10:
		fmt.Print("Nama Terlalu Panjang")
	case Panjang>5:
		fmt.Println("Panjang nama sudah cukup")
	default:
		fmt.Println("Nama sudah benar")
	}
}