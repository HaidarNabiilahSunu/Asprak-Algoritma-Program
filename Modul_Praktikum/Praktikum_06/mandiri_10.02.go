//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel a (waktu mulai), b (waktu selesai), dan durasi bertipe int
	var a, b, durasi int

	// Membaca input nilai waktu mulai (a) dan waktu selesai (b) dari pengguna
	fmt.Scan(&a, &b)

	// Mengecek apakah waktu selesai lebih besar atau sama dengan waktu mulai (berada dalam rentang hari yang sama)
	if b >= a {
		// Menghitung selisih durasi secara langsung (b dikurangi a)
		durasi = b - a
	} else {
		// Jika b < a, artinya waktu melewati pergantian hari (tengah malam / format 24 jam)
		// Menghitung durasi dengan sisa jam hari pertama (24 - a) ditambah jam pada hari berikutnya (b)
		durasi = (24 - a) + b
	}

	// Menampilkan hasil akhir total durasi waktu ke konsol
	fmt.Println(durasi)
}