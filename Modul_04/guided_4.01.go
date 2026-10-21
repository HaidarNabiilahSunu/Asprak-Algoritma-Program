//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan durasi jam, menit, dan detik
	var detik, jam, menit int

	// Membaca input total waktu dalam satuan detik dari pengguna
	fmt.Scan(&detik)

	// Menghitung jumlah jam (1 jam = 3600 detik)
	jam = detik / 3600

	// Menghitung jumlah menit dari sisa detik setelah diambil nilai jam (1 menit = 60 detik)
	menit = (detik % 3600) / 60

	// Menghitung sisa detik akhir setelah dikurangi nilai jam dan menit
	detik = detik % 60

	// Menampilkan hasil konversi ke format kalimat (X jam Y menit dan Z detik)
	fmt.Println(jam, "jam", menit, "menit dan", detik, "detik")
}