//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n (jumlah data boolean) dan j (counter perulangan)
	var n, j int
	// Deklarasi variabel nilai (untuk menampung input boolean per iterasi) dan hasil (akumulator logika)
	var nilai, hasil bool

	// Membaca input jumlah data boolean (n) dari pengguna
	fmt.Scan(&n)
	
	// Inisialisasi awal nilai hasil dengan true sebagai elemen identitas untuk operasi logika AND
	hasil = true

	// Perulangan dari j = 1 sampai j = n untuk membaca dan mengevaluasi setiap nilai boolean
	for j = 1; j <= n; j += 1 {
		// Membaca input nilai boolean (true/false) untuk iterasi saat ini
		fmt.Scan(&nilai)
		
		// Melakukan operasi logika AND (&&) antara akumulasi hasil sebelumnya dengan nilai input baru
		hasil = hasil && nilai
	}

	// Menampilkan hasil akhir dari evaluasi logika AND seluruh masukan
	fmt.Println(hasil)
}