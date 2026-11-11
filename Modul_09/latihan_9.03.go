//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel x dan y bertipe int untuk menampung dua nilai masukan bilangan bulat
	var x, y int
	// Deklarasi variabel hasil1 dan hasil2 bertipe bool untuk menampung status kebenaran kondisi kelipatan/faktor
	var hasil1, hasil2 bool

	// Membaca input nilai x dan y secara berurutan dari pengguna
	fmt.Scan(&x, &y)

	// Mengecek apakah y habis dibagi x (artinya y merupakan kelipatan dari x, atau x adalah faktor dari y)
	hasil1 = y % x == 0
	// Mengecek apakah x habis dibagi y (artinya x merupakan kelipatan dari y, atau y adalah faktor dari x)
	hasil2 = x % y == 0

	// Menampilkan hasil evaluasi apakah y kelipatan dari x (true atau false)
	fmt.Println(hasil1)
	// Menampilkan hasil evaluasi apakah x kelipatan dari y (true atau false)
	fmt.Println(hasil2)
}