//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi konstanta untuk harga satu cangkir kopi (Rp20.000) dan tarif pajak (11%)
	const hargaKopi int = 20000
	const pajak float64 = 0.11

	// Deklarasi variabel untuk jumlah pesanan kopi dan total pembayaran
	var jumlahKopi int
	var totalHarga float64

	// Membaca input jumlah cangkir kopi yang dibeli dari pengguna
	fmt.Scan(&jumlahKopi)

	// Menghitung subtotal harga sebelum pajak (dengan konversi tipe data ke float64)
	totalHarga = float64(jumlahKopi * hargaKopi)

	// Menghitung total harga akhir setelah ditambahkan pajak sebesar 11%
	totalHarga = totalHarga + (totalHarga * pajak)

	// Menampilkan hasil akhir total pembayaran yang dibulatkan tanpa angka di belakang koma
	fmt.Printf("%.0f", totalHarga)
}