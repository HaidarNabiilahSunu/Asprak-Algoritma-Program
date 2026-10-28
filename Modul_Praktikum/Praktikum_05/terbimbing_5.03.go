//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel N (batas akhir/jumlah bilangan) dan j (counter perulangan)
	var N, j int

	// Membaca input nilai N dari pengguna
	fmt.Scan(&N)

	// Perulangan dari j = 1 sampai j = N untuk menghasilkan deret kuadrat bilangan
	for j = 1; j <= N; j += 1 {
		// Menghitung hasil kuadrat dari j (j * j) dan mencetaknya diikuti spasi pada baris yang sama
		fmt.Print(j*j, " ")
	}
}