//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah data), i (counter perulangan), sisi, luas, dan keliling persegi
	var n, i, sisi, luas, keliling int

	// Membaca input jumlah data/perulangan n dari pengguna
	fmt.Scan(&n)

	// Perulangan dari i = 1 sampai i = n untuk memproses setiap input panjang sisi
	for i = 1; i <= n; i++ {
		// Membaca input panjang sisi persegi untuk iterasi saat ini
		fmt.Scan(&sisi)

		// Menghitung luas persegi dengan rumus: sisi * sisi
		luas = sisi * sisi

		// Menghitung keliling persegi dengan rumus: 4 * sisi
		keliling = 4 * sisi

		// Menampilkan hasil perhitungan luas dan keliling yang dipisahkan oleh spasi
		fmt.Println(luas, keliling)
	}
}