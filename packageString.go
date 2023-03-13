package main

import (
	"fmt"
	"strings"
)

func main() {
	//melakukan pengecekan apakah di string rizki maulana ada string rizki
	fmt.Println(strings.Contains("Rizki Maulana", "Rizki"))
	fmt.Println(strings.Contains("Rizki Maulana", "rizki")) //hasilnya false karena huruf r nya bukan huruf besar jadi contains akan menyatakan hasil rizki tidak ada

	//untuk memperkecil huruf string
	fmt.Println(strings.ToLower("Rizki Maulana"))
	//untuk memperbesar huruf string
	fmt.Println(strings.ToUpper("Rizki Maulana"))
	fmt.Println(strings.ToTitle("Rizki Maulana"))

	//untuk memotong spasi yang ada dikanan dan kiri
	fmt.Println(strings.Trim("           Rizki Maulana             ", " "))
	fmt.Println(strings.Trim("a           Rizki Maulana             a", " ")) //nah hasilnya yang ini tidak akan memotong huruf a walaupun di
	fmt.Println(strings.Trim("kontol Rizki Maulana    ", "kontol"))

	//nah split ini dia akan menjadikan nama rizki maulana menjadi dua nama rizki dan maulana
	fmt.Println(strings.Split("Rizki Maulana", " "))

	fmt.Println(strings.ReplaceAll("Rizki Maulana", "Rizki Maulana", "Budi"))
}
