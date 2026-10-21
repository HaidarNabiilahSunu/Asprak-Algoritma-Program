//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah sesi/data), i (counter loop), serta variabel konversi waktu
	var n, i int
	var jam, menit, sisa, durasi int

	// Membaca input jumlah data atau sesi (n) dari pengguna
	fmt.Scan(&n)

	// Inisialisasi awal akumulator total durasi dalam satuan menit
	durasi = 0

	// Perulangan dari i = 1 sampai i = n untuk membaca dan mengakumulasi input menit
	for i = 1; i <= n; i++ {
		// Membaca masukan durasi menit untuk iterasi saat ini
		fmt.Scan(&menit)

		// Menambahkan nilai menit masukan ke total akumulasi durasi
		durasi = durasi + menit
	}

	// Menghitung jumlah jam dengan pembagian bulat (1 jam = 60 menit)
	jam = durasi / 60

	// Menghitung sisa menit menggunakan operasi sisa bagi (modulo 60)
	sisa = durasi % 60
	menit = sisa

	// Menampilkan hasil konversi total waktu dalam format jam dan sisa menit
	fmt.Println(jam, menit)
}