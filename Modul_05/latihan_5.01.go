//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah perulangan), i (counter perulangan), dan teks (kalimat/kata yang dicetak)
	var n, i int
	var teks string

	// Membaca input jumlah perulangan n dari pengguna
	fmt.Scan(&n)

	// Membaca input teks/kalimat yang akan dicetak dari pengguna
	fmt.Scan(&teks)

	// Perulangan dari i = 1 sampai i = n untuk mencetak teks sebanyak n kali
	for i = 1; i <= n; i++ {
		// Menampilkan teks pada baris baru di setiap iterasi
		fmt.Println(teks)
	}
}