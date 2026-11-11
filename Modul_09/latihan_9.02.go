//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel bilangan bertipe int untuk menampung input dari pengguna
	var bilangan int
	// Deklarasi variabel teks bertipe string untuk menyimpan hasil keterangan kondisi
	var teks string

	// Membaca input nilai bilangan dari pengguna
	fmt.Scan(&bilangan)

	// Memberikan nilai awal default "bukan" pada variabel teks
	teks = "bukan"

	// Mengecek apakah bilangan merupakan bilangan genap (habis dibagi 2) DAN bernilai negatif (kurang dari 0)
	if bilangan % 2 == 0 && bilangan < 0 {
		// Mengubah nilai teks menjadi "genap negatif" jika kondisi tersebut bernilai benar
		teks = "genap negatif"
	}

	// Menampilkan hasil akhir string keterangan ke konsol
	fmt.Println(teks)
}