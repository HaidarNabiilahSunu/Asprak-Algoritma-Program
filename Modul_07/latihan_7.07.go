//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel x (batas awal), y (batas akhir), j (counter perulangan), dan jumlah (akumulator hasil)
	var x, y, j, jumlah int

	// Membaca input nilai batas awal (x) dan batas akhir (y) dari pengguna
	fmt.Scan(&x, &y)

	// Inisialisasi awal nilai akumulator jumlah dengan 0 sebelum proses penjumlahan dimulai
	jumlah = 0
	// Menentukan bilangan genap pertama yang bernilai lebih besar atau sama dengan x sebagai titik awal
	j = x + (x % 2)

	// Perulangan dimulai dari j dengan increment 2 (melompati bilangan ganjil) sampai batas y
	for ; j <= y; j += 2 {
		// Menambahkan nilai j (bilangan genap) saat ini ke dalam akumulator jumlah pada setiap iterasi
		jumlah = jumlah + j
	}

	// Menampilkan hasil akhir total penjumlahan deret bilangan genap dari x hingga y
	fmt.Println(jumlah)
}