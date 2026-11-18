//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel bilangan (untuk angka 4 digit) dan d1 s.d. d4 (untuk menampung masing-masing digit)
	var bilangan, d1, d2, d3, d4 int
	// Deklarasi variabel teks bertipe string untuk menyimpan status keterangan urutan digit
	var teks string

	// Menampilkan pesan panduan/prompt ke konsol sebelum membaca input
	fmt.Print("Bilangan: ")
	// Membaca input nilai bilangan bulat dari pengguna
	fmt.Scan(&bilangan)

	// Mengekstrak digit satuan (posisi paling kanan) menggunakan operasi sisa bagi (modulo) 10
	d4 = bilangan % 10
	// Mengekstrak digit puluhan dengan mengambil sisa bagi 100 kemudian dibagi 10
	d3 = (bilangan % 100) / 10
	// Mengekstrak digit ratusan dengan mengambil sisa bagi 1000 kemudian dibagi 100
	d2 = (bilangan % 1000) / 100
	// Mengekstrak digit ribuan (posisi paling kiri) dengan membagi bilangan secara bulat dengan 1000
	d1 = bilangan / 1000

	// Mengecek apakah urutan digit dari kiri ke kanan semakin besar secara ketat (d1 < d2 < d3 < d4)
	if d1 < d2 && d2 < d3 && d3 < d4 {
		// Mengubah nilai teks menjadi "terurut membesar" jika kondisi terpenuhi
		teks = "terurut membesar"
	// Mengecek apakah urutan digit dari kiri ke kanan semakin kecil secara ketat (d1 > d2 > d3 > d4)
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		// Mengubah nilai teks menjadi "terurut mengecil" jika kondisi terpenuhi
		teks = "terurut mengecil"
	// Jika tidak memenuhi kedua kondisi di atas (acak atau ada digit yang sama)
	} else {
		// Mengubah nilai teks menjadi "tidak terurut"
		teks = "tidak terurut"
	}

	// Menampilkan hasil akhir berupa informasi bilangan beserta status urutan digitnya ke konsol
	fmt.Println("Digit pada bilangan", bilangan, teks)
}