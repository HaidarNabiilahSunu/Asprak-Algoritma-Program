//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel a, b, c (nilai masukan sensor 0 atau 1) dan variabel sensor bertipe boolean
	var a, b, c int
	var sensor bool

	// Membaca input nilai dari tiga sensor (a, b, dan c)
	fmt.Scan(&a, &b, &c)

	// Evaluasi kondisi boolean: bernilai true jika setidaknya 2 dari 3 sensor bernilai 1 
	// (sistem mayoritas/dua dari tiga)
	sensor = (a == 1 && b == 1) ||
		(a == 1 && c == 1) ||
		(b == 1 && c == 1)

	// Menampilkan hasil evaluasi status sensor (true jika mayoritas aktif, false jika tidak)
	fmt.Println(sensor)
}