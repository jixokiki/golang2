package main
import "fmt"
func main(){
	//init statement contoh => i := 1 (ini dilakukan hanya diawal saja)
	//post statement contoh => i = i + 1/ i++ (ini dilakukan untuk mengecheck secara berkala sebuah perulangan)

	for i :=1 ; i < 10; i++{
		fmt.Println("Urutan : ", i)
	}

	splice := []string{"Rizki", "Hito", "Owen"}
	for i:= 0; i < len(splice); i++{
		fmt.Println(splice[i])
	}

	for i, value:= range splice{
		fmt.Println("Urutan", i, "=", value)
	}

	splice1 := append(splice, "Jeanet")
	fmt.Println(splice1)

	for _, value:= range splice1{
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["Urutan 1 :"] = "Rizki"
	person["Urutan 2 :"] = "Hito"

	for key, data:= range person {
		fmt.Println(key, "=", data)
	}
}