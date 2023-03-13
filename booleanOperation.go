package main
import "fmt"
func main()  {
	NilaiAkhir := 90
	AbsensiAkhir := 80

	SyaratNilaiKelulusan := NilaiAkhir >= 80
	SyaratAbsensiKelulusan := AbsensiAkhir >= 80

	Lulus := SyaratNilaiKelulusan && SyaratAbsensiKelulusan
	fmt.Println(Lulus)

	// atau bisa langsung dinyatakan tanpa membuat parametenya lagi seperti berikut
	fmt.Println(NilaiAkhir >= 80 && AbsensiAkhir >= 80, "Selamat Anda Lulus")
	fmt.Println(NilaiAkhir >= 80 || AbsensiAkhir > 90, "Selamat Mengulang")
}