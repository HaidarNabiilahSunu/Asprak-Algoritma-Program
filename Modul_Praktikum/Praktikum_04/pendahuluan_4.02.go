//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah bulan/periode), i (counter loop), tabungan_awal, dan kenaikan
	var i, n, tabungan_awal, kenaikan int
	// Deklarasi variabel total (akumulator tabungan) dan penambahan (nilai tambahan tiap periode)
	var total, penambahan int

	// Membaca input nilai n (periode), tabungan_awal, dan kenaikan konstan dari pengguna
	fmt.Scan(&n, &tabungan_awal, &kenaikan)

	// Inisialisasi saldo awal total tabungan
	total = tabungan_awal

	// Inisialisasi nilai penambahan awal yang besarnya sama dengan kenaikan
	penambahan = kenaikan

	// Perulangan dari i = 1 sampai i = n untuk mensimulasikan penambahan tabungan bertingkat
	for i = 1; i <= n; i++ {
		// Menambahkan nilai penambahan periode saat ini ke total tabungan
		total = total + penambahan

		// Meningkatkan nilai penambahan untuk periode berikutnya sesuai besar kenaikan
		penambahan = penambahan + kenaikan
	}

	// Menampilkan hasil akhir total akumulasi tabungan setelah n periode
	fmt.Println(total)
}