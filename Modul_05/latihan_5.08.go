//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (angka masukan), i (counter perulangan), dan hasil bertipe int64 untuk menampung perkalian besar
	var n, i int
	var hasil int64

	// Membaca input nilai n dari pengguna untuk dihitung faktorialnya (n!)
	fmt.Scan(&n)

	// Inisialisasi awal nilai hasil dengan 1 sebagai identitas perkalian
	hasil = 1

	// Perulangan dari i = 1 sampai i = n untuk melakukan akumulasi perkalian faktorial
	for i = 1; i <= n; i++ {
		// Mengalikan nilai hasil saat ini dengan i (dengan konversi tipe data ke int64)
		hasil = hasil * int64(i)
	}

	// Menampilkan hasil akhir perhitungan faktorial
	fmt.Println(hasil)
}