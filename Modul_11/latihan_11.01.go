//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel ph bertipe float64 untuk menampung nilai derajat keasaman (pH) air
	var ph float64

	// Membaca input nilai desimal pH air dari pengguna
	fmt.Scan(&ph)

	// Percabangan switch (expression-less) untuk mengevaluasi kelayakan air berdasarkan tingkat pH
	switch {
	// Kondisi 1: Memvalidasi apakah nilai pH berada di luar rentang standar kimia (kurang dari 0 atau lebih dari 14)
	case ph < 0 || ph > 14:
		// Menampilkan pesan error apabila nilai pH tidak valid/di luar batas skala
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")

	// Kondisi 2: Mengecek apakah nilai pH berada di rentang normal atau netral (antara 6.5 sampai 8.6)
	case ph >= 6.5 && ph <= 8.6:
		// Menampilkan pesan bahwa air aman dan memenuhi syarat untuk dikonsumsi (layak minum)
		fmt.Println("Air layak minum")

	// Kondisi default: Dijalankan jika pH valid (antara 0-14) tetapi berada di luar rentang layak minum (terlalu asam atau terlalu basa)
	default:
		// Menampilkan pesan bahwa air tidak memenuhi syarat untuk dikonsumsi
		fmt.Println("Air tidak layak minum")
	}
}