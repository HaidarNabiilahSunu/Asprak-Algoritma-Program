//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel totalDetik (masukan waktu), jam, menit, dan detik bertipe int
	var totalDetik, jam, menit, detik int

	// Membaca input nilai total detik dari pengguna
	fmt.Scan(&totalDetik)

	// Menghitung jumlah jam (1 jam setara dengan 3600 detik)
	jam = totalDetik / 3600
	
	// Menghitung jumlah menit dari sisa detik setelah dikurangi jam (1 menit setara dengan 60 detik)
	menit = (totalDetik % 3600) / 60
	
	// Menghitung sisa detik akhir yang tidak cukup untuk membentuk 1 menit
	detik = totalDetik % 60

	// Menampilkan hasil konversi waktu dalam format jam, menit, dan detik
	fmt.Println(jam, menit, detik)
}