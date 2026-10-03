package main

import "fmt"

func main() {
	// Deklarasi variabel tahun bertipe int dan kabisat bertipe boolean (true/false)
	var tahun int
	var kabisat bool

	// Membaca input nilai tahun dari pengguna
	fmt.Scan(&tahun)

	// Mengecek apakah tahun merupakan tahun kabisat:
	// - Habis dibagi 400, ATAU
	// - Habis dibagi 4 TETAPI tidak habis dibagi 100
	kabisat = (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	// Menampilkan tahun yang diinputkan
	fmt.Printf("Tahun: %d\n", tahun)

	// Menampilkan status kabisat (true/false) menggunakan format %t
	fmt.Printf("Kabisat: %t\n", kabisat)
}