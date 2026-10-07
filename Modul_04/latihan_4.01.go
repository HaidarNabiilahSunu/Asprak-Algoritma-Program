//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel total belanja, persentase diskon, dan total pembayaran akhir
	var totalBelanja, diskon int
	var totalAkhir float64

	// Membaca input nilai total belanja dari pengguna
	fmt.Scan(&totalBelanja)

	// Membaca input persentase diskon dari pengguna
	fmt.Scan(&diskon)

	// Menghitung total pembayaran setelah dipotong diskon (dengan konversi tipe data ke float64)
	totalAkhir = float64(totalBelanja) - (float64(totalBelanja) * float64(diskon) / 100)

	// Menampilkan hasil total akhir pembayaran yang dikonversi kembali ke bilangan bulat (int)
	fmt.Println(int(totalAkhir))
}