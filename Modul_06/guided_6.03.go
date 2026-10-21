//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel j (counter loop), v1 (bilangan yang dikalikan), dan v2 (pengali/jumlah perulangan)
	var j, v1, v2 int
	// Deklarasi variabel hasil untuk menampung total penjumlahan berulang
	var hasil int

	// Membaca input nilai v1 dan v2 dari pengguna
	fmt.Scan(&v1, &v2)

	// Inisialisasi awal nilai hasil dengan 0
	hasil = 0

	// Perulangan sebanyak v2 kali untuk mensimulasikan perkalian melalui penjumlahan berulang (v1 + v1 + ... sebanyak v2 kali)
	for j = 1; j <= v2; j += 1 {
		// Menambahkan nilai v1 ke akumulator hasil pada setiap iterasi
		hasil = hasil + v1
	}

	// Menampilkan hasil akhir penjumlahan berulang (setara dengan v1 * v2)
	fmt.Println(hasil)
}