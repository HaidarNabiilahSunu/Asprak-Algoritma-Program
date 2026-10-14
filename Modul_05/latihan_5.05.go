//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah data), i (counter perulangan), dan bilangan (input angka 4 digit)
	var n, i, bilangan int
	// Deklarasi variabel untuk menyimpan digit pertama, digit terakhir, dan hasil penjumlahan
	var digitPertama, digitTerakhir, hasil int

	// Membaca input jumlah data/perulangan n dari pengguna
	fmt.Scan(&n)

	// Perulangan dari i = 1 sampai i = n untuk memproses setiap masukan bilangan
	for i = 1; i <= n; i++ {
		// Membaca input bilangan 4 digit untuk iterasi saat ini
		fmt.Scan(&bilangan)

		// Mengambil digit pertama (ribuan) dengan pembagian bulat 1000
		digitPertama = bilangan / 1000

		// Mengambil digit terakhir (satuan) menggunakan operasi modulo 10
		digitTerakhir = bilangan % 10

		// Menjumlahkan digit pertama dan digit terakhir
		hasil = digitPertama + digitTerakhir

		// Menampilkan hasil penjumlahan untuk iterasi saat ini
		fmt.Println(hasil)
	}
}