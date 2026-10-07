//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel N (urutan hari ke-N) dan hari (hasil indeks hari 1-7)
	var N, hari int

	// Membaca input jumlah hari N dari pengguna
	fmt.Scan(&N)

	// Menghitung hari ke-N menggunakan pola aritmatika modulo 7
	// Angka 4 menentukan hari awal (misal: Kamis jika minggu dimulai dari Senin = 1)
	hari = (4+N-1)%7 + 1

	// Menampilkan urutan hari hasil perhitungan (nilai antara 1 hingga 7)
	fmt.Println(hari)
}