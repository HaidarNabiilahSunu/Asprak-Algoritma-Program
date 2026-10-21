//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan jumlah koin dan konversinya (emas, perak, tembaga)
	var emas, sisa, perak, tembaga, koin int

	// Membaca input total koin dari pengguna
	fmt.Scan(&koin)

	// Menghitung jumlah koin emas (1 emas = 9 koin)
	emas = koin / 9

	// Menghitung sisa koin setelah dikonversi ke emas
	sisa = koin % 9

	// Menghitung jumlah koin perak dari sisa koin (1 perak = 3 koin)
	perak = sisa / 3

	// Menghitung sisa akhir koin yang menjadi koin tembaga
	tembaga = koin % 3

	// Menampilkan hasil penukaran koin (Emas Perak Tembaga)
	fmt.Println(emas, perak, tembaga)
}