//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel jumlahOrang (total orang) dan jumlahMotor (kebutuhan motor)
	var jumlahOrang, jumlahMotor int

	// Membaca input jumlah total orang dari pengguna
	fmt.Scan(&jumlahOrang)

	// Menghitung estimasi awal jumlah motor dengan asumsi 1 motor dapat dinaiki maksimal 2 orang
	jumlahMotor = jumlahOrang / 2

	// Mengecek apakah jumlah orang ganjil (sisa pembagian dengan 2 tidak sama dengan 0)
	if jumlahOrang % 2 != 0 {
		// Menambahkan 1 motor ekstra untuk menampung sisa orang yang tidak mendapat pasangan
		jumlahMotor = jumlahMotor + 1
	}

	// Menampilkan hasil akhir jumlah total motor minimum yang dibutuhkan
	fmt.Println(jumlahMotor)
}