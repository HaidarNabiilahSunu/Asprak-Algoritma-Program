//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel angka bertipe int untuk menampung nomor hari (rentang 1 sampai 7) dari pengguna
	var angka int

	// Menampilkan prompt/pesan panduan ke konsol sebelum meminta input angka
	fmt.Print("Masukkan angka (1-7): ")
	// Membaca input nilai bilangan bulat dari pengguna
	fmt.Scan(&angka)

	// Percabangan switch untuk mencocokkan nilai variabel angka dengan nama hari dalam seminggu
	switch angka {
	// Jika angka bernilai 1, mencetak nama hari Senin
	case 1:
		fmt.Println("Senin")
	// Jika angka bernilai 2, mencetak nama hari Selasa
	case 2:
		fmt.Println("Selasa")
	// Jika angka bernilai 3, mencetak nama hari Rabu
	case 3:
		fmt.Println("Rabu")
	// Jika angka bernilai 4, mencetak nama hari Kamis
	case 4:
		fmt.Println("Kamis")
	// Jika angka bernilai 5, mencetak nama hari Jumat
	case 5:
		fmt.Println("Jumat")
	// Jika angka bernilai 6, mencetak nama hari Sabtu
	case 6:
		fmt.Println("Sabtu")
	// Jika angka bernilai 7, mencetak nama hari Minggu
	case 7:
		fmt.Println("Minggu")
	// Blok default: dijalankan apabila angka yang dimasukkan di luar rentang 1 s.d. 7
	default:
		fmt.Println("Hari tidak valid")
	}
}