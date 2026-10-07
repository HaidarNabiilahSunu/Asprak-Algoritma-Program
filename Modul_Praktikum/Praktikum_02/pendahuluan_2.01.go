//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah uang/poin) bertipe int dan beli bertipe boolean (true/false)
	var n int
	var beli bool

	// Membaca input nilai n dari pengguna
	fmt.Scan(&n)

	// Mengecek kondisi apakah nilai n memenuhi syarat minimal (1500 atau lebih)
	beli = n >= 1500

	// Menampilkan status keputusan beli (true jika n >= 1500, false jika tidak)
	fmt.Println(beli)
}