//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel bilangan untuk menampung nilai input dari pengguna
	var bilangan int
	// Deklarasi variabel teks untuk menyimpan status keterangan jenis bilangan
	var teks string

	// Membaca input nilai bilangan dari pengguna
	fmt.Scan(&bilangan)

	// Memberikan nilai awal default "negatif (bukan positif)" pada variabel teks
	teks = "negatif (bukan positif)"

	// Mengecek apakah nilai bilangan lebih besar dari 0 (positif)
	if bilangan > 0 {
		// Mengubah nilai teks menjadi "positif" jika kondisi bilangan > 0 bernilai benar
		teks = "positif"
	}

	// Menampilkan hasil akhir string status keterangan bilangan ke konsol
	fmt.Println(teks)
}