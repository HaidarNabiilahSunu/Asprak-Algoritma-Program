//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (batas akhir penjumlahan) dan j (counter perulangan)
	var n, j int
	// Deklarasi variabel hasil untuk menampung akumulasi total penjumlahan
	var hasil int

	// Membaca input nilai n dari pengguna
	fmt.Scan(&n)

	// Inisialisasi awal nilai hasil dengan 0
	hasil = 0

	// Perulangan dari j = 1 sampai j = n untuk menjumlahkan setiap deret angka (1 + 2 + ... + n)
	for j = 1; j <= n; j += 1 {
		// Menambahkan nilai j saat ini ke akumulator hasil
		hasil = hasil + j
	}

	// Menampilkan hasil akhir akumulasi deret bilangan
	fmt.Println(hasil)
}