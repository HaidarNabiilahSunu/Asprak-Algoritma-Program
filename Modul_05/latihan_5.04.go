//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah subjek/orang), i dan j (counter perulangan), langkah (input harian), serta totalLangkah
	var n, i, j, langkah, totalLangkah int

	// Membaca input jumlah subjek atau orang (n) dari pengguna
	fmt.Scan(&n)

	// Perulangan luar: memproses data untuk setiap subjek dari ke-1 hingga ke-n
	for i = 1; i <= n; i++ {
		// Mengatur ulang akumulator total langkah menjadi 0 sebelum menghitung minggu baru
		totalLangkah = 0

		// Perulangan dalam: membaca dan menjumlahkan input langkah harian selama 7 hari (1 minggu)
		for j = 1; j <= 7; j++ {
			// Membaca input jumlah langkah untuk hari ke-j
			fmt.Scan(&langkah)

			// Menambahkan langkah harian ke akumulator total langkah minggu tersebut
			totalLangkah += langkah
		}

		// Menampilkan akumulasi total langkah mingguan untuk subjek ke-i
		fmt.Println(totalLangkah)
	}
}