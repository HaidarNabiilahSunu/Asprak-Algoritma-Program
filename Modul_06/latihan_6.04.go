//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (bilangan masukan) dan j (counter perulangan)
	var n, j int
	// Deklarasi variabel hasil untuk menampung akumulasi nilai faktorial
	var hasil int

	// Membaca input nilai n dari pengguna
	fmt.Scan(&n)

	// Inisialisasi awal nilai hasil dengan 1 sebagai identitas perkalian
	hasil = 1

	// Perulangan dari j = 1 sampai j = n untuk melakukan perkalian beruntun (1 * 2 * ... * n)
	for j = 1; j <= n; j += 1 {
		// Mengalikan nilai hasil saat ini dengan j pada setiap iterasi
		hasil = hasil * j
	}

	// Menampilkan hasil akhir perhitungan faktorial
	fmt.Println(hasil)
}