//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan 3-digit angka masukan dan masing-masing digitnya
	var bilangan, d1, d2, d3 int

	// Membaca input bilangan 3-digit dari pengguna
	fmt.Scan(&bilangan)

	// Mengambil digit pertama (ratusan)
	d1 = bilangan / 100

	// Mengambil digit kedua (puluhan)
	d2 = bilangan % 100 / 10

	// Mengambil digit ketiga (satuan)
	d3 = bilangan % 100 % 10

	// Menampilkan status boolean (true/false) apakah urutan digitnya naik/terurut (d1 <= d2 <= d3)
	fmt.Println(d1 <= d2 && d2 <= d3)
}