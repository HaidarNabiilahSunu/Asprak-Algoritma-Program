//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel daging (jenis bahan utama), menu (pilihan olahan), dan tambahan (bahan pelengkap) bertipe string
	var daging, menu, tambahan string

	// Membaca input jenis daging dan pilihan menu dari pengguna secara berurutan
	fmt.Scan(&daging, &menu)

	// Cabang luar (outer if): memeriksa apakah jenis daging yang dipilih adalah "ayam"
	if daging == "ayam" {
		// Cabang dalam (inner if): memeriksa jenis menu spesifik jika dagingnya adalah ayam
		if menu == "soup" {
			// Menetapkan bahan tambahan "sayur" jika menunya adalah soup
			tambahan = "sayur"
		} else if menu == "katsu" {
			// Menetapkan bahan tambahan "tepung" jika menunya adalah katsu
			tambahan = "tepung"
		}
	// Cabang luar alternatif (else if): memeriksa apakah jenis daging yang dipilih adalah "sapi"
	} else if daging == "sapi" {
		// Cabang dalam: memeriksa jenis menu spesifik jika dagingnya adalah sapi
		if menu == "teriyaki" {
			// Menetapkan bahan tambahan "bombay" jika menunya adalah teriyaki
			tambahan = "bombay"
		} else if menu == "semur" {
			// Menetapkan bahan tambahan "kecap" jika menunya adalah semur
			tambahan = "kecap"
		}
	}

	// Menampilkan informasi jenis daging yang dipilih ke konsol
	fmt.Println("Daging :", daging)
	// Menampilkan hasil penentuan bahan tambahan berdasarkan logika percabangan di atas
	fmt.Println("Tambahan :", tambahan)
	
}