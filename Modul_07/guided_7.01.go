//go:build ignore

// Versi - Menggunakan Perulangan
package main

import "fmt"

func main() {
	// Deklarasi variabel x (batas awal), y (batas akhir), j (counter perulangan), dan jumlah (akumulator hasil)
	var x, y, j, jumlah int

	// Membaca input nilai batas awal (x) dan batas akhir (y) dari pengguna
	fmt.Scan(&x, &y)

	// Inisialisasi awal nilai akumulator jumlah dengan 0 sebelum proses penjumlahan dimulai
	jumlah = 0

	// Perulangan dari j = x sampai j = y dengan inkremen 1 untuk menjumlahkan setiap bilangan dalam rentang tersebut
	for j = x; j <= y; j += 1 {
		// Menambahkan nilai j saat ini ke dalam akumulator jumlah pada setiap iterasi
		jumlah = jumlah + j
	}

	// Menampilkan hasil akhir total penjumlahan rentang bilangan dari x hingga y
	fmt.Println(jumlah)
}