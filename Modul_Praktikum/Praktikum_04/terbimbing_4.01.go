//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah data/iterasi), i (counter loop), dan h (penampung hasil perhitungan)
	var n, i, h int
	// Deklarasi variabel b1 hingga b6 untuk menyimpan 6 bilangan bulat masukan
	var b1, b2, b3, b4, b5, b6 int

	// Membaca input jumlah data atau perulangan (n) dari pengguna
	fmt.Scan(&n)

	// Perulangan dari i = 1 sampai i = n untuk memproses setiap kumpulan 6 bilangan
	for i = 1; i <= n; i++ {
		// Membaca 6 bilangan bulat sekaligus untuk iterasi saat ini
		fmt.Scan(&b1, &b2, &b3, &b4, &b5, &b6)

		// Menghitung ekspresi aritmatika berdasarkan urutan presedensi operator (*, /, % dieksekusi sebelum +, -)
		h = b1 + b2 - b3*b4/b5%b6

		// Menampilkan hasil perhitungan aritmatika untuk iterasi saat ini
		fmt.Println(h)
	}
}