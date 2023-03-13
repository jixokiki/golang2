package main
import "fmt"
func main(){
	//var nama = "Rizki" 
	var nama = "Hito"
	//var nama = "Owen"
	
	if nama == "Rizki"{ //jika yang dijadiin penentunya/var rizki 
		fmt.Println("Selamat datang Rizki")
	}else if nama == "Hito"{//jika yang dijadiin penentunya/var hito
		fmt.Println("Selamat datang Hito")
	}else if nama == "Owen"{//jika yang dijadiin penentunya/var owen
		fmt.Println("Selamat datang Owen")
	}else{
		fmt.Println("Silahkan konfirmasi terlebih dahulu")
	}

	//var length = len(nama)
	// didalam golang short statement diatas atau short statement baru bisa ditaruh didalam if statement dengan menambahkan ;
	if length:= len(nama);length>5{
		fmt.Println("Nama terlalu panjang")
	}else{
		fmt.Println("Nama sudah benar")
	}
}