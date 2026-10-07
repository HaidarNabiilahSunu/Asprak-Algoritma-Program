//go:build ignore
package main

import "fmt"

// Tipe bentukan untuk menyimpan data persegi panjang
type PersegiPanjang struct {
	panjang  float64
	lebar    float64
	luas     float64
	keliling float64
}

func main() {
	var p PersegiPanjang

	// Input panjang dan lebar
	fmt.Print("Panjang: ")
	fmt.Scan(&p.panjang)

	fmt.Print("Lebar: ")
	fmt.Scan(&p.lebar)

	// Menghitung luas persegi panjang
	p.luas = p.panjang * p.lebar

	// Menghitung keliling persegi panjang
	p.keliling = 2 * (p.panjang + p.lebar)

	// Menampilkan hasil
	fmt.Printf("Luas: %.1f\n", p.luas)
	fmt.Printf("Keliling: %.1f\n", p.keliling)
}