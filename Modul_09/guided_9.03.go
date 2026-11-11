//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel bilangan bertipe int untuk menampung input dari pengguna
	var bilangan int
	// Deklarasi variabel hasil bertipe bool untuk menampung status kebenaran dari kondisi
	var hasil bool

	// Membaca input nilai bilangan dari pengguna
	fmt.Scan(&bilangan)

	// Mengecek apakah bilangan merupakan bilangan genap (bilangan % 2 == 0) DAN bernilai negatif (bilangan < 0)
	hasil = bilangan % 2 == 0 && bilangan < 0

	// Menampilkan hasil akhir evaluasi logika boolean (true atau false) ke konsol
	fmt.Println(hasil)
}