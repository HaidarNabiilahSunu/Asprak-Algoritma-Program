//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel n bertipe int untuk menampung nilai bilangan bulat masukan dari pengguna
	var n int

	// Membaca input nilai bilangan bulat dari pengguna
	fmt.Scan(&n)

	// Percabangan switch (expression-less) untuk mengkategorikan bilangan berdasarkan aturan logika tertentu
	switch {
	// Kondisi 1: Jika n adalah 5 ATAU jika n adalah bilangan ganjil yang bukan kelipatan 5
	case n == 5 || (n%2 != 0 && n%5 != 0):
		// Menampilkan kategori bahwa bilangan tersebut masuk kelompok bilangan ganjil
		fmt.Println("Kategori: Bilangan Ganjil")
		// Menampilkan operasi matematika: penjumlahan n dengan bilangan berikutnya (n + 1)
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n",
			n, n+1, n+(n+1))

	// Kondisi 2: Jika n merupakan kelipatan 10 (habis dibagi 10)
	case n%10 == 0:
		// Menampilkan kategori bilangan kelipatan 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		// Menampilkan operasi matematika: pembagian nilai n dengan angka 10
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n",
			n, n/10)

	// Kondisi 3: Jika n merupakan kelipatan 5 (habis dibagi 5, namun bukan kelipatan 10 karena sudah disaring di atas)
	case n%5 == 0:
		// Menampilkan kategori bilangan kelipatan 5
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		// Menampilkan operasi matematika: penghitungan kuadrat dari n (n dikali n)
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n",
			n, n*n)

	// Kondisi 4: Jika n adalah bilangan genap (habis dibagi 2)
	case n%2 == 0:
		// Menampilkan kategori bilangan genap
		fmt.Println("Kategori: Bilangan Genap")
		// Menampilkan operasi matematika: perkalian n dengan bilangan berikutnya (n * (n + 1))
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n",
			n, n+1, n*(n+1))
	}
}