//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel kode bertipe string untuk menyimpan kode huruf kendaraan dari pengguna
	var kode string

	// Menampilkan prompt/pesan panduan ke konsol sebelum meminta input kode
	fmt.Print("Masukkan kode kendaraan (A-D): ")
	// Membaca input string berupa kode huruf dari pengguna
	fmt.Scan(&kode)

	// Percabangan switch untuk mencocokkan nilai variabel kode dengan jenis kendaraan yang sesuai
	switch kode {
	// Jika kode bernilai "A", menampilkan jenis kendaraan "Mobil"
	case "A":
		fmt.Println("Mobil")

	// Jika kode bernilai "B", menampilkan jenis kendaraan "Motor"
	case "B":
		fmt.Println("Motor")

	// Jika kode bernilai "C", menampilkan jenis kendaraan "Sepeda"
	case "C":
		fmt.Println("Sepeda")

	// Jika kode bernilai "D", menampilkan jenis kendaraan "Bus"
	case "D":
		fmt.Println("Bus")

	// Blok default: dijalankan apabila kode yang dimasukkan di luar pilihan A, B, C, atau D
	default:
		fmt.Println("Jenis Kendaraan Tidak Tersedia")
	}
}