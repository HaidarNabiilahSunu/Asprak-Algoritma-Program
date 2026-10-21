//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah iterasi/data), i (counter perulangan), x (input resep/adonan), dan totalTepung
	var n, i, x, totalTepung int

	// Membaca input jumlah data/perulangan n dari pengguna
	fmt.Scan(&n)

	// Perulangan sebanyak n kali untuk memproses setiap input masukan
	for i = 1; i <= n; i++ {
		// Membaca nilai x untuk iterasi saat ini
		fmt.Scan(&x)

		// Menghitung kebutuhan total tepung dengan mengalikan nilai x dengan faktor pengali 10
		totalTepung = x * 10

		// Menampilkan hasil perhitungan total tepung untuk iterasi saat ini
		fmt.Println(totalTepung)
	}
}