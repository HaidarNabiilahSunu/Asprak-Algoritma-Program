//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel s bertipe int untuk menampung nilai skor/poin 
	// dan variabel hasil untuk menyimpan keterangan kelas
	var s int
	var hasil string

	// Membaca input nilai s dari pengguna
	fmt.Scan(&s)

	// Mengecek apakah nilai s lebih besar dari atau sama dengan 450
	if s >= 450 {
		// Menetapkan hasil ke "KELAS REGULER" jika syarat nilai minimal terpenuhi
		hasil = "KELAS REGULER"
	} else {
		// Menetapkan hasil ke "KELAS MATRIKULASI" jika nilai berada di bawah batas reguler
		hasil = "KELAS MATRIKULASI"
	}

	// Menampilkan hasil penentuan kelas akhir ke konsol
	fmt.Println(hasil)
}