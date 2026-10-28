//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk satuan mata uang: dinar, dirham, fals, dan qirat
	var dinar, dirham, fals, qirat int

	// Membaca input total nilai dalam satuan qirat dari pengguna
	fmt.Scan(&qirat)

	// Menghitung jumlah dinar (asumsi 1 dinar = 600 qirat)
	dinar = qirat / 600

	// Menghitung jumlah dirham dari sisa qirat setelah dikurangi dinar (asumsi 1 dirham = 60 qirat)
	dirham = qirat % 600 / 60

	// Menghitung jumlah fals dari sisa qirat sebelumnya (asumsi 1 fals = 6 qirat)
	fals = qirat % 600 % 60 / 6

	// Menghitung sisa akhir dalam satuan qirat yang tidak cukup untuk dikonversi ke fals
	qirat = qirat % 600 % 60 % 6

	// Menampilkan hasil konversi satuan mata uang dalam urutan dinar, dirham, fals, dan qirat
	fmt.Println(dinar, dirham, fals, qirat)
}