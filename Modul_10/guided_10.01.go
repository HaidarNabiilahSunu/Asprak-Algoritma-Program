//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel usia (untuk menampung umur) dan kk (status kepemilikan Kartu Keluarga bertipe boolean)
	var usia int
	var kk bool

	// Membaca input nilai usia dan status kepemilikan KK (true/false) secara berurutan dari pengguna
	fmt.Scan(&usia, &kk)

	// Mengecek apakah usia sudah 17 tahun atau lebih DAN status kepemilikan KK bernilai true (terpenuhi)
	if usia >= 17 && kk {
		// Menampilkan pesan jika syarat usia dan Kartu Keluarga terpenuhi
		fmt.Println("bisa membuat KTP")
	} else {
		// Menampilkan pesan jika salah satu atau kedua syarat di atas tidak terpenuhi
		fmt.Println("belum bisa membuat KTP")
	}
}