package main
import "fmt"

func main(){
	var name [4] string
	name[0] = "Rizki"
	name[1] = "Hito"
	name[2] = "Jeanet"
	name[3] = "Diva"

	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	fmt.Println(name[3])

	var Kumpulan = [3]string{
		"Rizki",
		"Jeanet",
		"Diva",
	}

	fmt.Println(Kumpulan)

	fmt.Println(len (name))
	fmt.Println(len(Kumpulan))

	fmt.Println(len(name[1]))

	var nama [10] string

	fmt.Println(len(nama))

	Kumpulan[0] = "Wahid Bani"
	fmt.Println(Kumpulan)
}