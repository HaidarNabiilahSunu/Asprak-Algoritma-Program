//go:build ignore
package main

import "fmt"

// Tipe bentukan untuk menyimpan data transaksi
type Transaksi struct {
	nama       	string
	jumlah     	int
	hargaSatuan float64
	totalHarga 	float64
}

func main() {
	var t Transaksi

	fmt.Println("Informasi Transaksi")

	// Input data transaksi
	fmt.Print("Nama Barang : ")
	fmt.Scan(&t.nama)

	fmt.Print("Jumlah : ")
	fmt.Scan(&t.jumlah)

	fmt.Print("Harga Satuan : ")
	fmt.Scan(&t.hargaSatuan)

	// Menghitung total harga barang
	t.totalHarga = float64(t.jumlah) * t.hargaSatuan

	// Menampilkan hasil transaksi
	fmt.Printf("Total Harga: Rp %.2f\n", t.totalHarga)
}