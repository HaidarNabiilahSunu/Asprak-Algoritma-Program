//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel s (saldo) dan p (harga pembayaran) bertipe int
	var s, p int

	// Membaca input nilai saldo (s) dan harga pembayaran (p) dari pengguna
	fmt.Scan(&s, &p)

	// Validasi batasan agar nilai s dan p berada dalam rentang 0 sampai 1.000.000
	if (s >= 0 && s <= 1000000) && (p >= 0 && p <= 1000000) {
		// Mengecek apakah saldo lebih besar atau sama dengan pembayaran
		if s >= p {
			// Menampilkan pesan jika saldo mencukupi 
			// (Catatan: kondisi s == p secara logika sudah tertangani dan tercakup di dalam blok s >= p ini)
			fmt.Print("PEMBAYARAN BERHASIL")
		} else if s == p {
			// Cabang ini secara logika tidak akan pernah tercapai (dead code) karena sudah disaring oleh kondisi s >= p di atasnya
			fmt.Print("SALDO SANGAT PAS")
		} else {
			// Menampilkan pesan jika saldo lebih kecil dari harga pembayaran (tidak cukup)
			fmt.Print("SALDO TIDAK CUKUP")
		}
	}
}