//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah data/hari), i (counter perulangan), jam (input durasi), dan totalJam (akumulator)
	var n, i, jam, totalJam int
	// Deklarasi variabel rataRata bertipe float64 untuk menyimpan hasil pembagian desimal
	var rataRata float64

	// Membaca input jumlah data atau durasi pengamatan (n) dari pengguna
	fmt.Scan(&n)

	// Inisialisasi awal nilai totalJam menjadi 0 sebelum penjumlahan dimulai
	totalJam = 0

	// Perulangan dari i = 1 sampai i = n untuk membaca dan menjumlahkan durasi jam
	for i = 1; i <= n; i++ {
		// Membaca input jumlah jam untuk iterasi saat ini
		fmt.Scan(&jam)

		// Menambahkan nilai jam masukan ke akumulator totalJam
		totalJam += jam
	}

	// Menghitung nilai rata-rata dengan mengonversi totalJam dan n ke float64 agar menghasilkan pembagian presisi/desimal
	rataRata = float64(totalJam) / float64(n)

	// Menampilkan hasil rata-rata dengan format 3 angka di belakang koma
	fmt.Printf("%.3f\n", rataRata)
}