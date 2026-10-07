//go:build ignore
package main

import "fmt"

// Tipe bentukan untuk menyimpan data BMI seseorang
type BMI struct {
	nama   string
	berat  float64
	tinggi float64
	bmi    float64
}

func main() {
	var b BMI

	// Input data seseorang
	fmt.Print("Nama: ")
	fmt.Scan(&b.nama)

	fmt.Print("Berat: ")
	fmt.Scan(&b.berat)

	fmt.Print("Tinggi: ")
	fmt.Scan(&b.tinggi)

	// Menghitung BMI
	b.bmi = b.berat / (b.tinggi * b.tinggi)

	// Menampilkan nilai BMI
	fmt.Printf("BMI: %.2f\n", b.bmi)
}