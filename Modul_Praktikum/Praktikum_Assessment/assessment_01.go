//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel jumlahKategori untuk menampung total kategori yang akan diinput
	var jumlahKategori int

	// Meminta dan membaca input jumlah kategori dari pengguna
	fmt.Print("Masukkan jumlah kategori : ")
	fmt.Scan(&jumlahKategori)

	// Perulangan luar (outer loop) untuk mengiterasi setiap kategori dari 1 sampai jumlahKategori
	for i := 1; i <= jumlahKategori; i++ {
		// Deklarasi variabel jumlahBarang untuk menampung total barang dalam kategori saat ini
		var jumlahBarang int

		// Menampilkan informasi nomor kategori saat ini
		fmt.Printf("\nKategori ke-%d\n", i)
		fmt.Print("Masukkan jumlah barang : ")
		fmt.Scan(&jumlahBarang)

		// Perulangan dalam (inner loop) untuk mengiterasi setiap barang dari 1 sampai jumlahBarang dalam kategori tersebut
		for j := 1; j <= jumlahBarang; j++ {
			// Deklarasi variabel untuk menyimpan data detail barang (nama, harga, stok)
			var nama string
			var harga, stok int

			// Menampilkan informasi nomor barang saat ini
			fmt.Printf("\nBarang ke-%d\n", j)

			// Meminta dan membaca input nama barang
			fmt.Print("Nama Barang : ")
			fmt.Scan(&nama)

			// Meminta dan membaca input harga barang
			fmt.Print("Harga Barang : ")
			fmt.Scan(&harga)

			// Meminta dan membaca input stok barang
			fmt.Print("Stok Barang : ")
			fmt.Scan(&stok)

			// Menampilkan kembali data barang yang baru saja dimasukkan ke konsol
			fmt.Println("Nama :", nama)
			fmt.Println("Harga :", harga)
			fmt.Println("Stok :", stok)
		}
	}
}