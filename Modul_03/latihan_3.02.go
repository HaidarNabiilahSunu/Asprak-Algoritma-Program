//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel jejari (r) bertipe int, serta volume dan luas bertipe float64
	var r int
	var volume, luas float64

	// Deklarasi konstanta nilai Pi
	const pi = 3.1415926535

	// Membaca input nilai jejari (r) dari pengguna
	fmt.Scan(&r)

	// Menghitung volume bola V = (4/3) * pi * r^3
	// float64() digunakan untuk konversi tipe r agar perkalian presisi
	volume = (4.0 / 3.0) * pi * float64(r*r*r)

	// Menghitung luas permukaan bola L = 4 * pi * r^2
	luas = 4 * pi * float64(r*r)

	// Menampilkan nilai jejari yang diinputkan
	fmt.Printf("Jejari = %d\n", r)

	// Menampilkan hasil perhitungan volume dan luas dengan format 4 angka di belakang koma
	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n",
		r, volume, luas)
}