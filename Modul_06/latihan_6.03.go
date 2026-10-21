//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel bilangan (basis), pangkat (eksponen), j (counter loop), dan hasil
	var bilangan, pangkat int
	var j int
	var hasil int

	// Membaca input nilai bilangan basis dan nilai pangkat dari pengguna
	fmt.Scan(&bilangan, &pangkat)

	// Inisialisasi awal nilai hasil dengan 1 sebagai identitas perkalian
	hasil = 1

	// Perulangan sebanyak nilai pangkat untuk mengalikan bilangan basis secara berulang
	for j = 1; j <= pangkat; j += 1 {
		// Mengalikan nilai hasil saat ini dengan bilangan basis pada setiap iterasi
		hasil = hasil * bilangan
	}

	// Menampilkan hasil akhir perpangkatan
	fmt.Println(hasil)
}