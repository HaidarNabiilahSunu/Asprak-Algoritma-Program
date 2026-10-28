//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel a, b, c, dan d bertipe float64 untuk menampung elemen-elemen matriks berukuran 2x2
	var a, b, c, d float64
	// Deklarasi variabel hasil bertipe float64 untuk menampung nilai determinan
	var hasil float64

	// Membaca input nilai keempat elemen matriks (a, b, c, d) secara berurutan dari pengguna
	fmt.Scan(&a, &b, &c, &d)

	// Menghitung nilai determinan matriks 2x2 menggunakan rumus: (a * d) - (b * c)
	hasil = (a * d) - (b * c)

	// Menampilkan hasil akhir perhitungan determinan matriks
	fmt.Println(hasil)
}