package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a])o")
	fmt.Println(regex.MatchString("Iki"))
	fmt.Println(regex.MatchString("iki"))
	fmt.Println(regex.MatchString("ikiu"))

	search := regex.FindAllString("iki iki kiki kuki koki", -1)
	fmt.Println(search)
}
