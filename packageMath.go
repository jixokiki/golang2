package main

import (
	"fmt"
	"math"
)

func main() {
	//ini membulatkan kebawah
	fmt.Println(math.Round(1.2))

	//ini membulatkan keatas karena melebihi dari minimalnya yaitu 1.5
	fmt.Println(math.Round(1.7))

	//ini memaksa membulatkan keatas
	fmt.Println(math.Ceil(1.2))
	fmt.Println(math.Ceil(0.1))

	//ini memaksa membulatkan kebawah
	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Floor(1.2))

	//menyatakan nilai maximal
	fmt.Println(math.Max(10, 30))

	//menyatakan nilai minimal
	fmt.Println(math.Min(100, 20))
}
