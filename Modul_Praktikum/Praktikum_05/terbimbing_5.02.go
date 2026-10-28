//go:build ignore

// Versi - Menggunakan Rumus Deret Aritmatika
package main

import "fmt"

func main() {
	// Deklarasi variabel x (batas awal), y (batas akhir), dan n (jumlah suku dalam rentang)
	var x, y, n int
	// Deklarasi variabel S bertipe float64 untuk menampung hasil perhitungan jumlah deret (sum)
	var S float64

	// Membaca input nilai batas awal (x) dan batas akhir (y) dari pengguna
	fmt.Scan(&x, &y)

	// Menghitung total banyaknya suku/angka (n) dalam rentang dari x sampai y
	n = (y - x) + 1

	// Menghitung jumlah total deret menggunakan rumus aritmatika Gauss: S = (n * (x + y)) / 2
	S = float64(n*(x+y)) / 2

	// Menampilkan hasil akhir perhitungan jumlah total deret bilangan
	fmt.Println(S)
}