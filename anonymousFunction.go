package main

import "fmt"

type Blacklist func(string) bool

func userIn(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Wellcome ", name)
	}
}

func main() {
	blacklists := func(name string) bool {
		return name == "admin"
	}

	userIn("admin", blacklists)
	userIn("rizki", blacklists)

	//atau bisa langsung panggil
	userIn("iki", func(name string) bool {
		return name == "admin"
	})
	userIn("admin", func(name string) bool {
		return name == "admin"
	})
}
