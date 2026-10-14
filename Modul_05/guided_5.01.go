//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (batas awal), m (batas akhir), dan i (variabel counter perulangan)
	var n, m, i int

	// Membaca input nilai n (awal) dan m (akhir) dari pengguna
	fmt.Scan(&n, &m)

	// Perulangan untuk mencetak angka secara berurutan mulai dari n hingga m
	for i = n; i <= m; i++ {
		// Menampilkan nilai i diikuti dengan spasi
		fmt.Printf("%d ", i)
	}
	// Mencetak baris baru (newline) setelah perulangan selesai
	fmt.Println()
}
