package main

import "fmt"

func main() {
	// Deklarasi variabel r (jari-jari) dan luas dengan tipe data float64
	var r, luas float64

	// Membaca input nilai jari-jari dari pengguna
	fmt.Scan(&r)

	// Menghitung luas lingkaran menggunakan rumus L = π * r^2
	luas = 3.14 * r * r

	// Menampilkan hasil perhitungan luas dengan format 1 angka di belakang koma
	fmt.Printf("%.1f\n", luas)
}