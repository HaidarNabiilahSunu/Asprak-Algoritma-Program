//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel a (batas awal) dan b (batas akhir)
	var a, b int
	// Deklarasi variabel j sebagai counter perulangan
	var j int

	// Membaca input nilai rentang a dan b dari pengguna
	fmt.Scan(&a, &b)

	// Perulangan dari j = a sampai j = b dengan inkremen 2
	for j = a; j <= b; j += 2 {
		// Menampilkan nilai j saat ini diikuti oleh spasi pada baris yang sama
		fmt.Print(j, " ")
	}
}