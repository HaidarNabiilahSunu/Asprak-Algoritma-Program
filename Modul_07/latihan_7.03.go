//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk koordinat titik pertama (x1, y1, z1) bertipe float64
	var x1, y1, z1 float64
	// Deklarasi variabel untuk koordinat titik kedua (x2, y2, z2) bertipe float64
	var x2, y2, z2 float64
	// Deklarasi variabel untuk koordinat titik ketiga (x3, y3, z3) bertipe float64
	var x3, y3, z3 float64
	// Deklarasi variabel untuk menampung hasil rata-rata masing-masing sumbu (rataX, rataY, rataZ)
	var rataX, rataY, rataZ float64

	// Membaca input nilai koordinat titik pertama dari pengguna
	fmt.Scan(&x1, &y1, &z1)
	// Membaca input nilai koordinat titik kedua dari pengguna
	fmt.Scan(&x2, &y2, &z2)
	// Membaca input nilai koordinat titik ketiga dari pengguna
	fmt.Scan(&x3, &y3, &z3)

	// Menghitung nilai rata-rata untuk sumbu X dari ketiga titik
	rataX = (x1 + x2 + x3) / 3
	// Menghitung nilai rata-rata untuk sumbu Y dari ketiga titik
	rataY = (y1 + y2 + y3) / 3
	// Menghitung nilai rata-rata untuk sumbu Z dari ketiga titik
	rataZ = (z1 + z2 + z3) / 3

	// Menampilkan hasil akhir rata-rata koordinat (rataX, rataY, rataZ)
	fmt.Println(rataX, rataY, rataZ)
}