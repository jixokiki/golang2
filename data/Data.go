package data

var koneksi string

func init() {
	koneksi = "Uwampp"
}

func GetData() string {
	return koneksi
}
