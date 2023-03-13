package main

import (
	"fmt"
	"os"
)

func main() {
	// untuk mendeklarasikan atau membuat argument dari terminal
	result := os.Args
	fmt.Println("Arguments or location file :")
	fmt.Println(result)

	fmt.Println(result[1])
	fmt.Println(result[2])

	//meliat hostname atau pengecekan eror hostname
	namaHost, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", namaHost)
	} else {
		fmt.Println("Terjadi Eror : ", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
