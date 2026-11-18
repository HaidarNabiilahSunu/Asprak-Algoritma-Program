//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel b (bilangan yang dicari faktornya) dan i (variabel iterator/pencacah untuk perulangan) bertipe int
	var b, i int

	// Menampilkan pesan teks prompt ke konsol sebelum meminta input bilangan
	fmt.Print("Bilangan: ")
	// Membaca input nilai bilangan bulat dari pengguna
	fmt.Scan(&b)

	// Menampilkan teks label awal baris untuk daftar faktor
	fmt.Print("Faktor:")

	// Melakukan perulangan (for-loop) dari i = 1 sampai dengan nilai b
	for i = 1; i <= b; i++ {
		// Mengecek apakah bilangan b habis dibagi oleh i (sisa bagi / modulo bernilai 0)
		if b%i == 0 {
			// Jika habis dibagi, maka i merupakan faktor dari b, lalu cetak nilainya dengan spasi di depannya
			fmt.Print(" ", i)
		}
	}

	// Mencetak baris baru (newline) di akhir program agar tampilan terminal rapi
	fmt.Println()
}