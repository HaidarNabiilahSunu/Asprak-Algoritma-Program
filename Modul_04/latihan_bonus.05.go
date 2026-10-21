//go:build ignore
package main

import "fmt"

// Tipe bentukan untuk menyimpan data karyawan
type Karyawan struct {
	nama       string
	gajiPokok  float64
	tunjangan  float64
	potongan   float64
	totalGaji  float64
}

func main() {
	var k Karyawan

	// Input data karyawan
	fmt.Print("Nama: ")
	fmt.Scan(&k.nama)

	fmt.Print("Gaji Pokok: ")
	fmt.Scan(&k.gajiPokok)

	fmt.Print("Tunjangan: ")
	fmt.Scan(&k.tunjangan)

	fmt.Print("Potongan: ")
	fmt.Scan(&k.potongan)

	// Menghitung total gaji
	k.totalGaji = k.gajiPokok + k.tunjangan - k.potongan

	// Menampilkan total gaji
	fmt.Printf("Total Gaji: Rp %.2f\n", k.totalGaji)
}