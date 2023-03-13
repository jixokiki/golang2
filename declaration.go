package main
import "fmt"
func main()  {
	type NoNim string
	type Kehadiran bool

	var KTM NoNim = "20200801041"
	var Absen Kehadiran = true
	fmt.Println(KTM)
	fmt.Println(Absen)
}