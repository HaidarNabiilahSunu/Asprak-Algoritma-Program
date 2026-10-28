//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel panjang, lebar, tinggi, volume, dan luasPermukaan bertipe float64 untuk presisi desimal
	var panjang, lebar, tinggi float64
	var volume, luasPermukaan float64

	// Membaca input nilai panjang, lebar, dan tinggi balok secara berurutan dari pengguna
	fmt.Scan(&panjang, &lebar, &tinggi)

	// Menghitung volume balok dengan rumus: panjang * lebar * tinggi
	volume = panjang * lebar * tinggi
	
	// Menghitung luas permukaan balok dengan rumus: 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
	luasPermukaan = 2 * (panjang*lebar +
		panjang*tinggi + lebar*tinggi)

	// Menampilkan hasil akhir perhitungan volume dan luas permukaan balok
	fmt.Println(volume, luasPermukaan)
}