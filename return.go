package main
import "fmt"
//return value 
func absensi(NameFirst string)  string{
	//penerapan return yang dasar\\
	//==========================\\
	//result := "Hello"
	//return result

	//penerapan return yang biasa dipakai\\
	//==========================\\
	//return "Hello " + NameFirst

	//penerapan if else return yang dasar\\
	if NameFirst == ""|| NameFirst != "Iki"{
		return "Acces Denied"
	} else {
		return "Selamat Datang " + NameFirst
	}
}

//multiple return value
func Biodata() (string, int, string, int) {
	return "Rizki Maulana", 24, "September", 2002
}
func main()  {
	fmt.Println(absensi("Iki"))

	result := absensi("")
	fmt.Println(result)
	
	Nama, Tanggal, Bulan, Tahun  := Biodata()
	fmt.Println(Nama, Tanggal, Bulan, Tahun)
}