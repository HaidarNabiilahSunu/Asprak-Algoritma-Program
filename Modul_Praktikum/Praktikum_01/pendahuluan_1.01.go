//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel N (total detik) serta hasil konversi: Jam, Menit, dan Detik
	var N, Sisa, Jam, Menit, Detik int

	// Membaca input total waktu dalam satuan detik dari pengguna
	fmt.Scan(&N)

	// Menghitung jumlah jam (1 jam = 3600 detik)
	Jam = N / 3600

	// Menghitung sisa detik setelah diambil nilai jam
	Sisa = N % 3600

	// Menghitung jumlah menit dari sisa detik (1 menit = 60 detik)
	Menit = Sisa / 60

	// Menghitung sisa akhir yang menjadi satuan detik
	Detik = Sisa % 60

	// Menampilkan hasil konversi ke format (Jam Menit Detik)
	fmt.Println(Jam, Menit, Detik)
}