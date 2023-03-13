package main

import "fmt"

/**
fungsi nill digunakan untuk melakukan pengecekan
*/
func MapRuangan(ruangan string) map[string]string {
	if ruangan == "" {
		return nil
	} else {
		return map[string]string{
			"ruangan": ruangan,
		}
	}
}

func main() {
	//sebenernya deklarasi nill bisa secara langsung
	var desa map[string]string = nil
	fmt.Println(desa)

	kelas := MapRuangan("201")
	fmt.Println(kelas)

	//jika pengecekan tanpa cara nil
	var kelass map[string]string

	if kelass["201"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(kelass)
	}

	//jika pengecekan menggunakan cara nil
	var kelazz map[string]string = MapRuangan("202")
	if kelazz == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(kelazz)
	}
}
