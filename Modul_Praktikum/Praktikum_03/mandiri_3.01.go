//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan harga awal dan harga akhir setelah diskon
	var hargaAwal, hargaAkhir float64

	// Membaca input harga awal barang dari pengguna
	fmt.Scan(&hargaAwal)

	// Menghitung harga akhir dengan diskon bertingkat 
	// (diskon pertama 10% -> dikali 0.90, diskon kedua 20% -> dikali 0.80)
	hargaAkhir = hargaAwal * 0.90 * 0.80

	// Menampilkan hasil harga akhir dengan format 2 angka di belakang koma
	fmt.Printf("%.2f\n", hargaAkhir)
}