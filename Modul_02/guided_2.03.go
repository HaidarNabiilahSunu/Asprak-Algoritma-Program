//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel c1-c5 bertipe byte (untuk karakter) dan b1-b3 bertipe int
	var c1, c2, c3, c4, c5 byte
	var b1, b2, b3 int

	// Membaca 5 karakter sekaligus dari input pengguna
	fmt.Scan(&c1, &c2, &c3, &c4, &c5)

	// Membaca 3 karakter berikutnya satu per satu menggunakan format karakter (%c)
	fmt.Scanf("%c", &b1)
	fmt.Scanf("%c", &b2)
	fmt.Scanf("%cc", &b3)

	// Menampilkan kembali 5 karakter pertama sesuai input asli
	fmt.Printf("%c%c%c%c%c", c1, c2, c3, c4, c5)

	// Menampilkan 3 karakter berikutnya yang nilainya digeser 1 langkah ke karakter setelahnya (ASCII + 1)
	fmt.Printf("%c%c%c", b1+1, b2+1, b3+1)
}