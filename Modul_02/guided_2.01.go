//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi 5 variabel bilangan bulat untuk input dan 1 variabel untuk hasil
	var a, b, c, d, e int
	var hasil int

	// Membaca 5 input bilangan bulat sekaligus dalam satu baris dari pengguna
	fmt.Scanln(&a, &b, &c, &d, &e)

	// Menghitung total penjumlahan dari kelima bilangan
	hasil = a + b + c + d + e

	// Menampilkan nilai tiap variabel beserta hasil akhir penjelasannya
	fmt.Println("Hasil penjumlahan", a, b, c, d, e, "adalah", hasil)
}