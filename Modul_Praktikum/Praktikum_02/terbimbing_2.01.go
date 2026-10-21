//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel kedalaman (h), batas tekanan, dan hasil tekanan bertipe float64
	var h, Batas, Tekanan float64

	// Deklarasi konstanta massa jenis air laut (p = 1025 kg/m³) dan percepatan gravitasi (g = 9.8 m/s²)
	const p float64 = 1025.0
	const g float64 = 9.8

	// Membaca input kedalaman (h) dan batas maksimum tekanan dari pengguna
	fmt.Scan(&h, &Batas)

	// Menghitung tekanan hidrostatis menggunakan rumus P = p * g * h
	Tekanan = p * g * h

	// Menampilkan nilai tekanan hidrostatis yang didapat
	fmt.Println(Tekanan)

	// Menampilkan status boolean (true/false) apakah tekanan masih di bawah atau sama dengan batas aman
	fmt.Println(Tekanan <= Batas)
}