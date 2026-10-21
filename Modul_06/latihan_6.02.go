//go:build ignore
package main

import (
	"fmt"
	"math" // Mengimpor paket math untuk menggunakan konstanta math.Pi
)

func main() {
	// Deklarasi variabel j (counter loop) dan n (jumlah data kerucut yang dihitung)
	var j, n int
	// Deklarasi variabel r (jari-jari), t (tinggi), dan volume bertipe float64 untuk presisi desimal
	var r, t, volume float64

	// Membaca input jumlah data kerucut (n) dari pengguna
	fmt.Scan(&n)

	// Perulangan dari j = 1 sampai j = n untuk memproses setiap kerucut
	for j = 1; j <= n; j += 1 {
		// Membaca nilai jari-jari (r) dan tinggi (t) kerucut untuk iterasi saat ini
		fmt.Scan(&r, &t)

		// Menghitung volume kerucut dengan rumus: (π * r² * t) / 3
		volume = (math.Pi * r * r * t) / 3

		// Menampilkan hasil perhitungan volume kerucut untuk iterasi saat ini
		fmt.Println(volume)
	}
}