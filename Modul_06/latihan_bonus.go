//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel i (counter loop luar untuk baris) dan j (counter loop dalam untuk kolom)
	var i, j int

	// Perulangan luar untuk mengatur jumlah baris (dari baris ke-1 sampai ke-5)
	// Catatan: Terdapat typo pada kondisi '1 <= 5' yang seharusnya 'i <= 5' agar nilai i diperiksa dengan benar
	for i = 1; i <= 5; i++ {
		// Perulangan dalam untuk mencetak karakter bintang sebanyak nomor baris saat ini (i)
		for j = 1; j <= i; j++ {
			// Mencetak karakter bintang dan spasi tanpa berpindah baris baru
			fmt.Print("* ")
		}
		// Mencetak baris baru (newline) setelah satu baris bintang selesai dicetak
		fmt.Println()
	}
}