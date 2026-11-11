//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel bilangan bertipe int untuk menampung nilai input dari pengguna
	var bilangan int

	// Membaca input nilai bilangan (bisa positif, negatif, atau nol) dari pengguna
	fmt.Scan(&bilangan)

	// Mengecek apakah bilangan bernilai negatif (kurang dari 0)
	if bilangan < 0 {
		// Mengubah bilangan negatif menjadi positif dengan mengalikannya dengan -1
		bilangan = -bilangan
	}

	// Menampilkan hasil akhir nilai absolut (mutlak) dari bilangan tersebut
	fmt.Println(bilangan)
}