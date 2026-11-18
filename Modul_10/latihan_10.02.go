//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel nam (nilai akhir mata kuliah) bertipe float64 dan nmk (nilai mutu huruf) bertipe string
	var nam float64
	var nmk string

	// Menampilkan teks/prompt ke konsol sebelum meminta input nilai angka dari pengguna
	fmt.Print("Nilai akhir mata kuliah: ")
	// Membaca input nilai akhir berupa bilangan desimal (float64) dari pengguna
	fmt.Scan(&nam)

	// Memeriksa rentang nilai numerik untuk menentukan huruf mutu menggunakan tangga percabangan (if-else ladder)
	if nam > 80 {
		// Jika nilai lebih dari 80, menetapkan huruf mutu ke "A"
		nmk = "A"
	} else if nam > 72.5 {
		// Jika nilai di atas 72.5 hingga 80, menetapkan huruf mutu ke "AB"
		nmk = "AB"
	} else if nam > 65 {
		// Jika nilai di atas 65 hingga 72.5, menetapkan huruf mutu ke "B"
		nmk = "B"
	} else if nam > 57.5 {
		// Jika nilai di atas 57.5 hingga 65, menetapkan huruf mutu ke "BC"
		nmk = "BC"
	} else if nam > 50 {
		// Jika nilai di atas 50 hingga 57.5, menetapkan huruf mutu ke "C"
		nmk = "C"
	} else if nam > 40 {
		// Jika nilai di atas 40 hingga 50, menetapkan huruf mutu ke "D"
		nmk = "D"
	} else {
		// Jika nilai 40 atau kurang dari itu, menetapkan huruf mutu ke "E"
		nmk = "E"
	}

	// Menampilkan hasil konversi nilai akhir menjadi huruf mutu mata kuliah ke konsol
	fmt.Println("Nilai mata kuliah:", nmk)
}