//go:build ignore
package main

import "fmt"

// Deklarasi konstanta PI untuk perhitungan geometri lingkaran/tabung
const PI float64 = 3.14

// Definisi struct Tabung untuk menyimpan properti dan hasil perhitungan tabung
type Tabung struct {
	tinggi      int
	jari2       int
	luas        float64
	volume      float64
	luasAlas    float64
	luasDinding float64
}

func main() {
	// Deklarasi variabel t bertipe struct Tabung
	var t Tabung

	// Membaca input nilai jari-jari dan tinggi tabung dari pengguna
	fmt.Scan(&t.jari2, &t.tinggi)

	// Menghitung luas alas lingkaran (L = PI * r^2) dengan konversi tipe data ke float64
	t.luasAlas = PI * float64(t.jari2*t.jari2)

	// Menghitung luas selimut/dinding tabung (Ld = t * 2 * PI * r)
	t.luasDinding = float64(t.tinggi) * (2 * PI * float64(t.jari2))

	// Menghitung total luas permukaan tabung (2 * Luas Alas + Luas Dinding)
	t.luas = 2*t.luasAlas + t.luasDinding

	// Menghitung volume tabung (V = Luas Alas * tinggi)
	t.volume = t.luasAlas * float64(t.tinggi)

	// Menampilkan hasil luas permukaan tabung dengan format 2 angka di belakang koma
	fmt.Printf("%.2f\n", t.luas)

	// Menampilkan hasil volume tabung dengan format 2 angka di belakang koma
	fmt.Printf("%.2f\n", t.volume)
}