//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel c1 dan c2 untuk menyimpan karakter (byte), serta khusus1 dan khusus2 bertipe boolean
	var c1, c2 byte
	var khusus1, khusus2 bool

	// Membaca input dua karakter berturut-turut dari pengguna
	fmt.Scanf("%c%c", &c1, &c2)

	// Mengecek apakah c1 adalah karakter khusus (bukan huruf kapital, huruf kecil, maupun angka 0-9)
	khusus1 = !((c1 >= 'A' && c1 <= 'Z') ||
		(c1 >= 'a' && c1 <= 'z') ||
		(c1 >= '0' && c1 <= '9'))

	// Mengecek apakah c2 adalah karakter khusus (bukan huruf kapital, huruf kecil, maupun angka 0-9)
	khusus2 = !((c2 >= 'A' && c2 <= 'Z') ||
		(c2 >= 'a' && c2 <= 'z') ||
		(c2 >= '0' && c2 <= '9'))

	// Menampilkan status boolean karakter pertama (true jika karakter khusus, false jika alfanumerik)
	fmt.Println(khusus1)

	// Menampilkan status boolean karakter kedua (true jika karakter khusus, false jika alfanumerik)
	fmt.Println(khusus2)
}