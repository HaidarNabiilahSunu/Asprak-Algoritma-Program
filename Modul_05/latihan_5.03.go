//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah iterasi), i (counter perulangan), bilangan (input angka), dan hasil
	var n, i, bilangan, hasil int

	// Membaca input jumlah data/perulangan n dari pengguna
	fmt.Scan(&n)

	// Perulangan dari i = 1 sampai i = n untuk memproses setiap input angka
	for i = 1; i <= n; i++ {
		// Membaca input bilangan untuk iterasi saat ini
		fmt.Scan(&bilangan)

		// Menghitung hasil perkalian antara bilangan masukan dengan urutan iterasi ke-i
		hasil = bilangan * i

		// Menampilkan hasil perkalian untuk iterasi saat ini
		fmt.Println(hasil)
	}
}