//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel j (counter loop), alas, tinggi, dan n (jumlah segitiga yang dihitung)
	var j, alas, tinggi, n int
	// Deklarasi variabel luas bertipe float64 untuk menampung hasil perhitungan desimal
	var luas float64

	// Membaca input jumlah data segitiga (n) dari pengguna
	fmt.Scan(&n)

	// Perulangan dari j = 1 sampai j = n untuk memproses setiap segitiga
	for j = 1; j <= n; j += 1 {
		// Membaca nilai alas dan tinggi segitiga untuk iterasi saat ini
		fmt.Scan(&alas, &tinggi)

		// Menghitung luas segitiga dengan rumus 0.5 * (alas * tinggi) serta konversi tipe ke float64
		luas = 0.5 * float64(alas*tinggi)

		// Menampilkan hasil perhitungan luas segitiga untuk iterasi saat ini
		fmt.Println(luas)
	}
}