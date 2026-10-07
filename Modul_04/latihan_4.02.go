//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk target BMI, tinggi badan (m), dan hasil perhitungan berat badan (kg)
	var bmi, tinggiBadan, beratBadan float64

	// Membaca input nilai target BMI dari pengguna
	fmt.Scan(&bmi)

	// Membaca input nilai tinggi badan dari pengguna
	fmt.Scan(&tinggiBadan)

	// Menghitung berat badan yang dibutuhkan berdasarkan rumus: beratBadan = BMI * (tinggiBadan^2)
	beratBadan = bmi * (tinggiBadan * tinggiBadan)

	// Menampilkan hasil perhitungan berat badan yang dibulatkan tanpa angka di belakang koma
	fmt.Printf("%.0f", beratBadan)
}