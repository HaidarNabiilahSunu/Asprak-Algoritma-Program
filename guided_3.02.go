package main

import "fmt"

func main() {
	// Deklarasi variabel alas dan tinggi (integer), serta luas (float64)
	var alas, tinggi int
	var luas float64

	// Menampilkan pesan dan membaca input untuk alas segitiga
	fmt.Printf("Masukan Alas : ")
	fmt.Scan(&alas)

	// Menampilkan pesan dan membaca input untuk tinggi segitiga
	fmt.Printf("Masukan Tinggi : ")
	fmt.Scan(&tinggi)

	// Menghitung luas segitiga (L = 0.5 * alas * tinggi)
	// float64() digunakan untuk konversi tipe data agar sesuai dengan variabel luas
	luas = 0.5 * float64(alas*tinggi)

	// Menampilkan hasil perhitungan luas
	fmt.Println("Jawabanya = %.2f\n", luas)
}