package main

import "fmt"

func main() {
	// Deklarasi variabel rupiah dengan tipe data int
	var rupiah int

	// Menampilkan pesan dan membaca input jumlah uang dalam rupiah
	fmt.Printf("Masukan Uang : ")
	fmt.Scan(&rupiah)

	// Deklarasi variabel dolar dengan tipe data float64
	var dolar float64

	// Konversi nilai rupiah ke USD dengan asumsi kurs 1 USD = Rp 15.000
	// float64() digunakan untuk mengubah nilai rupiah ke float agar hasil pembagian presisi
	dolar = float64(rupiah) / 15000

	// Menampilkan hasil konversi ke USD dengan format 2 angka di belakang koma
	fmt.Printf("USD %.2f", dolar)
}