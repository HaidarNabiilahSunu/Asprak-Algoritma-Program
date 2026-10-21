//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel sisi dan volume dengan tipe data float64
	var sisi, volume float64

	// Menampilkan teks permintaan input ke pengguna
	fmt.Printf("Masukan Sisi : ")
	// Membaca input nilai panjang sisi dari pengguna
	fmt.Scan(&sisi)

	// Menghitung volume kubus menggunakan rumus V = s^3 (sisi * sisi * sisi)
	volume = sisi * sisi * sisi

	// Menampilkan hasil perhitungan volume dengan format 2 angka di belakang koma
	fmt.Printf("Jawabanya = %.2f\n", volume)
}