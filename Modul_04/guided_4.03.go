//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk berat badan (kg), tinggi badan (m), dan hasil BMI bertipe float64
	var beratBadan, tinggiBadan, bmi float64

	// Membaca input berat badan dan tinggi badan dari pengguna
	fmt.Scan(&beratBadan, &tinggiBadan)

	// Menghitung indeks massa tubuh dengan rumus BMI = beratBadan / (tinggiBadan^2)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)

	// Menampilkan hasil perhitungan BMI dengan format 2 angka di belakang koma
	fmt.Printf("%.2f", bmi)
}