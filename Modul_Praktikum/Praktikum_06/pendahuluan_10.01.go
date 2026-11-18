//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel j (untuk menyimpan durasi jam) dan tarif (untuk menyimpan biaya) bertipe int
	var j, tarif int

	// Membaca input nilai durasi jam dari pengguna
	fmt.Scan(&j)

	// Menetapkan nilai awal tarif dasar sebesar 10000
	tarif = 10000

	// Mengecek apakah durasi (j) lebih besar dari 5 jam
	if j > 5 {
		// Menambahkan biaya ekstra sebesar 4000 jika durasi melebihi 5 jam
		tarif += 4000
	}

	// Menampilkan hasil akhir total tarif pembayaran ke konsol
	fmt.Println(tarif)
}