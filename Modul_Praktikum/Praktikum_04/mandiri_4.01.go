//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah bilangan ganjil yang dicetak), i (counter loop), dan ganjil (nilai bilangan ganjil)
	var n, i, ganjil int

	// Membaca input jumlah bilangan ganjil (n) yang ingin ditampilkan
	fmt.Scan(&n)

	// Inisialisasi bilangan ganjil pertama dengan nilai 1
	ganjil = 1

	// Perulangan sebanyak n kali untuk mencetak deret bilangan ganjil
	for i = 1; i <= n; i++ {
		// Menampilkan nilai bilangan ganjil saat ini
		fmt.Println(ganjil)

		// Menambahkan nilai 2 untuk mendapatkan bilangan ganjil berikutnya
		ganjil = ganjil + 2
	}
}