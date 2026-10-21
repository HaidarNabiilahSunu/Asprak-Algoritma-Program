//go:build ignore
package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi variabel untuk koordinat titik A, B, dan C (x, y)
	var xA, yA, xB, yB, xC, yC float64
	// Deklarasi variabel untuk menyimpan panjang tiap sisi dan sisi terpanjang
	var sisiAB, sisiBC, sisiAC float64
	var sisiTerpanjang float64

	// Membaca input koordinat untuk masing-masing titik A, B, dan C
	fmt.Scan(&xA, &yA)
	fmt.Scan(&xB, &yB)
	fmt.Scan(&xC, &yC)

	// Menghitung jarak antar titik (panjang sisi) menggunakan rumus Euclidean Distance: √((x2-x1)² + (y2-y1)²)
	sisiAB = math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
	sisiBC = math.Sqrt(math.Pow(xC-xB, 2) + math.Pow(yC-yB, 2))
	sisiAC = math.Sqrt(math.Pow(xC-xA, 2) + math.Pow(yC-yA, 2))

	// Inisialisasi awal, mengasumsikan sisi AB sebagai sisi terpanjang
	sisiTerpanjang = sisiAB

	// Membandingkan dan memperbarui nilai jika sisi BC lebih panjang
	if sisiBC > sisiTerpanjang {
		sisiTerpanjang = sisiBC
	}

	// Membandingkan dan memperbarui nilai jika sisi AC lebih panjang
	if sisiAC > sisiTerpanjang {
		sisiTerpanjang = sisiAC
	}

	// Mengecek apakah nilai sisi terpanjang merupakan bilangan bulat (tanpa desimal)
	if sisiTerpanjang == math.Trunc(sisiTerpanjang) {
		// Menampilkan hasil tanpa desimal jika bernilai bulat
		fmt.Printf("%.0f", sisiTerpanjang)
	} else {
		// Menampilkan hasil dengan 2 angka di belakang koma jika bernilai pecahan
		fmt.Printf("%.2f", sisiTerpanjang)
	}
}