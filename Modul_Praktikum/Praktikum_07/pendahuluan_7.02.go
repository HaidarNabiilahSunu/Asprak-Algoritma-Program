//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n bertipe int untuk menampung nilai angka mahasiswa yang dimasukkan
	var n int

	// Menampilkan teks prompt/panduan ke konsol sebelum meminta input nilai
	fmt.Print("Masukkan nilai mahasiswa: ")
	// Membaca input nilai bilangan bulat dari pengguna
	fmt.Scan(&n)

	// Percabangan switch (expression-less) untuk menentukan predikat huruf berdasarkan rentang nilai angka
	switch {
	// Kondisi 1: Jika nilai berada di rentang 85 sampai 100
	case n >= 85 && n <= 100:
		fmt.Println("Predikat A")

	// Kondisi 2: Jika nilai berada di rentang 75 hingga kurang dari 85
	case n >= 75 && n < 85:
		fmt.Println("Predikat B")

	// Kondisi 3: Jika nilai berada di rentang 65 hingga kurang dari 75
	case n >= 65 && n < 75:
		fmt.Println("Predikat C")

	// Kondisi 4: Jika nilai berada di rentang 50 hingga kurang dari 65
	case n >= 50 && n < 65:
		fmt.Println("Predikat D")

	// Kondisi 5: Jika nilai berada di rentang 0 hingga kurang dari 50
	case n >= 0 && n < 50:
		fmt.Println("Predikat E")

	// Blok default: dijalankan apabila nilai yang dimasukkan di luar rentang valid (misal: negatif atau di atas 100)
	default:
		fmt.Println("Nilai tidak valid")
	}
}