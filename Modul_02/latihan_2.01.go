package main

import "fmt"

func main() {
	// Deklarasi variabel string untuk menyimpan input dan variabel sementara
	var (
		satu, dua, tiga string
		temp            string
	)

	// Meminta dan membaca input string pertama dari pengguna
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)

	// Meminta dan membaca input string kedua dari pengguna
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	// Meminta dan membaca input string ketiga dari pengguna
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	// Menampilkan urutan string sebelum dirotasi
	fmt.Println("Output awal =", satu, dua, tiga)

	// Proses pergeseran/rotasi nilai variabel menggunakan variabel temp:
	// satu -> mendapat nilai dua
	// dua   -> mendapat nilai tiga
	// tiga  -> mendapat nilai satu (yang disimpan di temp)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	// Menampilkan urutan string setelah nilai-nilainya digeser
	fmt.Println("Output akhir =", satu, dua, tiga)
}