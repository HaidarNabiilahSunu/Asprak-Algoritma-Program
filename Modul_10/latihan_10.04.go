//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel b (bilangan yang diuji), i (pencacah perulangan), dan jumlah (penghitung total faktor) bertipe int
	var b, i, jumlah int
	// Deklarasi variabel prima bertipe bool untuk menyimpan status apakah bilangan tersebut adalah bilangan prima
	var prima bool

	// Menampilkan pesan teks prompt ke konsol sebelum meminta input bilangan dari pengguna
	fmt.Print("Bilangan: ")
	// Membaca input nilai bilangan bulat dari pengguna
	fmt.Scan(&b)

	// Inisialisasi awal penghitung jumlah faktor dengan nilai 0
	jumlah = 0

	// Menampilkan teks label awal untuk daftar faktor
	fmt.Print("Faktor:")

	// Melakukan perulangan (for-loop) dari i = 1 sampai dengan nilai b untuk memeriksa setiap angka pembagi
	for i = 1; i <= b; i++ {
		// Mengecek apakah bilangan b habis dibagi oleh i (sisa bagi / modulo bernilai 0)
		if b%i == 0 {
			// Jika habis dibagi, cetak nilai i sebagai bagian dari faktor bilangan tersebut
			fmt.Print(" ", i)
			// Menambahkan 1 ke variabel jumlah setiap kali ditemukan faktor baru
			jumlah = jumlah + 1
		}
	}

	// Menentukan status prima: bernilai true jika jumlah total faktor tepat sama dengan 2 (yaitu 1 dan bilangan itu sendiri)
	prima = jumlah == 2

	// Mencetak baris baru (newline) agar format terminal rapi setelah mencetak daftar faktor
	fmt.Println()
	// Menampilkan hasil status evaluasi apakah bilangan tersebut adalah bilangan prima (true atau false)
	fmt.Println("Prima:", prima)
}