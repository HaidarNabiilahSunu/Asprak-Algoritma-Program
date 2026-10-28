//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah data benda) dan j (counter perulangan)
	var n, j int
	// Deklarasi variabel massaJenis, volume, dan massa bertipe int untuk perhitungan tiap benda
	var massaJenis, volume, massa int
	// Deklarasi variabel totalMassa (akumulator) dan rataRata bertipe float64 untuk presisi desimal
	var totalMassa, rataRata float64

	// Membaca input jumlah data benda (n) dari pengguna
	fmt.Scan(&n)
	
	// Inisialisasi awal nilai totalMassa dengan 0 sebelum akumulasi dimulai
	totalMassa = 0

	// Perulangan dari j = 1 sampai j = n untuk memproses setiap data benda
	for j = 1; j <= n; j += 1 {
		// Membaca input nilai massa jenis dan volume untuk benda pada iterasi saat ini
		fmt.Scan(&massaJenis, &volume)
		
		// Menghitung massa benda dengan rumus: massa jenis * volume
		massa = massaJenis * volume
		
		// Menambahkan massa benda saat ini ke dalam akumulator totalMassa (dengan konversi tipe ke float64)
		totalMassa = totalMassa + float64(massa)
	}

	// Menghitung nilai rata-rata massa seluruh benda (totalMassa dibagi dengan jumlah data n)
	rataRata = totalMassa / float64(n)
	
	// Menampilkan hasil akhir perhitungan rata-rata massa
	fmt.Println(rataRata)
}