//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel kendaraan bertipe string, serta durasi dan tarif bertipe int
	var kendaraan string
	var durasi int
	var tarif int

	// Menampilkan prompt ke konsol dan membaca input jenis kendaraan dari pengguna
	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)

	// Menampilkan prompt ke konsol dan membaca input durasi parkir dalam satuan jam
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	// Percabangan switch (expression-less) untuk menentukan besaran tarif berdasarkan jenis kendaraan dan durasi parkir
	switch {
	// Kondisi jika kendaraan adalah "Motor" dengan durasi antara 1 sampai 2 jam
	case kendaraan == "Motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000

	// Kondisi jika kendaraan adalah "Motor" dengan durasi lebih dari 2 jam
	case kendaraan == "Motor" && durasi > 2:
		tarif = 9000

	// Kondisi jika kendaraan adalah "Mobil" dengan durasi antara 1 sampai 2 jam
	case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000

	// Kondisi jika kendaraan adalah "Mobil" dengan durasi lebih dari 2 jam
	case kendaraan == "Mobil" && durasi > 2:
		tarif = 20000

	// Kondisi jika kendaraan adalah "Truk" dengan durasi antara 1 sampai 2 jam
	case kendaraan == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000

	// Kondisi jika kendaraan adalah "Truk" dengan durasi lebih dari 2 jam
	case kendaraan == "Truk" && durasi > 2:
		tarif = 35000

	// Blok default: dijalankan apabila input jenis kendaraan atau rentang durasi tidak sesuai dengan case di atas
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}

	// Menampilkan hasil akhir rincian nominal tarif parkir menggunakan format string (fmt.Printf)
	fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}