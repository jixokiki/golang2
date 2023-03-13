package main
import "fmt"
func main(){
	person := map[string]string{
		"Nama":"Rizki",
		"Usia":"20",
	}

	person["Gender"]="L"
	fmt.Println(person)
	fmt.Println(person["Nama"])
	fmt.Println(person["Usia"])

	var book map[string]string = make(map[string]string)
	book["Genre 1"] = "Python Full Course"
	book["Genre 2"] = "Golang Full Course"
	book["Genre 3"] = "Javascript Full Course"
	
	book["Genre 4"]="Construct Full Course"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "Genre 4")
	fmt.Println(book)
	fmt.Println(len(book))
}