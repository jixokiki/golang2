package main
import "fmt"
func main(){
	const fullName = "Rizki Maulana"
	const firstName string = "Rizki"
	const lastName = "Maulana"
	const age = 21 // klo tidak di print tidak akan terjadi eror karena tipe data nya constant dan data ini tidak bisa diubah lagi
	fmt.Println(fullName)
	fmt.Println(firstName)
	fmt.Println(lastName)

	//pemanggilan banyak dan simple dalam constant
	const (
		FN = "RIZKI MAULANA"
		Fst = "Rizki"
		Lst = "Maulana"
		U = 21 
	)

	fmt.Println(FN)
	fmt.Println(Fst)
	fmt.Println(Lst)
	fmt.Println(U)
}