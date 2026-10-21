//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel N (bilangan bulat masukan), d (pembagi/counter loop), dan s (status apakah d pembagi dari N)
	var N, d int
	var s bool

	// Membaca input nilai N dari pengguna
	fmt.Scan(&N)

	// Perulangan dari d = 1 sampai d = N untuk mengecek setiap calon pembagi
	for d = 1; d <= N; d++ {
		// Memeriksa apakah N habis dibagi d (sisa pembagian adalah 0)
		if N%d == 0 {
			// Jika habis dibagi, maka d adalah faktor/pembagi dari N (status true)
			s = true
		} else {
			// Jika tidak habis dibagi, maka d bukan faktor dari N (status false)
			s = false
		}

		// Menampilkan nilai d beserta status boolean s untuk iterasi saat ini
		fmt.Println(d, s)
	}
}