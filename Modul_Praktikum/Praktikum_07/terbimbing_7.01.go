//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel kode (pilihan menu), jumlah (banyak tiket), harga (satuan), dan total (keseluruhan pembayaran) bertipe int
	var kode, jumlah int
	// Deklarasi variabel genre bertipe string untuk menyimpan nama genre film yang dipilih
	var genre string
	var harga, total int

	// Menampilkan judul menu bioskop ke konsol
	fmt.Println("=== BIOSKOP ===")
	// Menampilkan pilihan menu genre film 1 (Action) dan 2 (Comedy)
	fmt.Println("1. Action")
	fmt.Println("2. Comedy")

	// Menampilkan prompt dan membaca input kode pilihan genre dari pengguna
	fmt.Print("Pilih genre film: ")
	fmt.Scan(&kode)

	// Menampilkan prompt dan membaca input jumlah tiket yang ingin dibeli
	fmt.Print("Jumlah tiket: ")
	fmt.Scan(&jumlah)

	// Percabangan switch untuk menentukan nama genre dan harga satuan tiket berdasarkan kode pilihan
	switch kode {
	// Jika pengguna memilih kode 1
	case 1:
		genre = "Action"
		harga = 50000

	// Jika pengguna memilih kode 2
	case 2:
		genre = "Comedy"
		harga = 40000

	// Blok default: dijalankan jika kode pilihan di luar angka 1 atau 2
	default:
		// Menampilkan pesan error pilihan tidak tersedia
		fmt.Println("Pilihan tidak tersedia")
		// Menghentikan eksekusi program lebih lanjut menggunakan perintah return
		return
	}

	// Memeriksa validitas jumlah tiket (harus lebih besar dari 0)
	if jumlah <= 0 {
		// Menampilkan pesan error jika jumlah tiket nol atau bernilai negatif
		fmt.Println("Jumlah tiket tidak valid")
		// Menghentikan eksekusi program
		return
	}

	// Menghitung total pembayaran dengan mengalikan harga satuan tiket dengan jumlah tiket yang dibeli
	total = harga * jumlah

	// Menampilkan rincian detail pemesanan tiket bioskop ke konsol
	fmt.Println("\n=== DETAIL PEMESANAN ===")
	fmt.Println("Genre:", genre)
	fmt.Println("Harga:", harga)
	fmt.Println("Jumlah tiket:", jumlah)
	fmt.Println("Total pembayaran:", total)
}