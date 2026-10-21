//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah iterasi/data), j (counter perulangan), x (input jarak dasar), dan totalJarak
	var n, j, x, totalJarak int

	// Membaca input jumlah data/perulangan n dari pengguna
	fmt.Scan(&n)

	// Perulangan sebanyak n kali untuk memproses setiap input masukan
	for j = 1; j <= n; j++ {
		// Membaca nilai x untuk iterasi saat ini
		fmt.Scan(&x)

		// Menghitung total jarak dengan mengalikan nilai x dengan faktor pengali 5
		totalJarak = x * 5

		// Menampilkan hasil perhitungan total jarak untuk iterasi saat ini
		fmt.Println(totalJarak)
	}
}