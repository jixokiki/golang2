package main
import "fmt"

func Penjumlahan(A int, B int){
	hasil := A + B
	fmt.Println("Hasil dari operasi penjumlahan berikut :", hasil)
}

func Introduce(firstName string, lastName string)  {
	fmt.Println("Selamat datang saudara ", firstName, lastName)
}
func main(){
	Penjumlahan(10,10)
	Introduce("Rizki", "Maulana")
}