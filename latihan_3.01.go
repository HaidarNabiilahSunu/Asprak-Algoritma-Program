package main

import "fmt"

func main() {
	// Deklarasi variabel x (input) dan fx (hasil fungsi) bertipe float64
	var x, fx float64

	// Membaca input nilai x dari pengguna
	fmt.Scan(&x)

	// Menghitung nilai f(x) berdasarkan rumus f(x) = (2 / (x + 5)) + 5
	fx = (2 / (x + 5)) + 5

	// Menampilkan hasil perhitungan f(x) ke layar
	fmt.Println(fx)
}