//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah request), i (counter loop), dan prioritas (nilai prioritas request)
	var n, i, prioritas int

	// Membaca input jumlah request (n) dari pengguna
	fmt.Scan(&n)

	// Inisialisasi tingkat prioritas awal sama dengan jumlah total request (prioritas tertinggi diawali dari angka terbesar)
	prioritas = n

	// Perulangan dari i = 1 sampai i = n untuk memproses setiap request berurutan
	for i = 1; i <= n; i++ {
		// Menampilkan pesan nomor request beserta tingkat prioritasnya untuk iterasi saat ini
		fmt.Println("Request", i,
			"diproses dengan prioritas", prioritas)

		// Menurunkan nilai prioritas sebesar 1 untuk request berikutnya
		prioritas = prioritas - 1
	}
}