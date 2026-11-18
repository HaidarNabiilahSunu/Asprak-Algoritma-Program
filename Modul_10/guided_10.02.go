//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel x bertipe rune untuk menampung satu karakter input dari pengguna
	var x rune
	// Deklarasi variabel boolean untuk menyimpan status kategori karakter (apakah huruf, vokal kecil, atau vokal besar)
	var huruf, vKecil, vBesar bool

	// Membaca input satu karakter tunggal dari pengguna menggunakan format '%c'
	fmt.Scanf("%c", &x)

	// Memeriksa apakah karakter x berada di dalam rentang alfabet (huruf kecil 'a'-'z' atau huruf besar 'A'-'Z')
	huruf = (x >= 'a' && x <= 'z') ||
		(x >= 'A' && x <= 'Z')

	// Memeriksa apakah karakter x merupakan salah satu huruf vokal kecil ('a', 'i', 'u', 'e', 'o')
	vKecil = x == 'a' || x == 'i' ||
		x == 'u' || x == 'e' || x == 'o'

	// Memeriksa apakah karakter x merupakan salah satu huruf vokal besar ('A', 'I', 'U', 'E', 'O')
	vBesar = x == 'A' || x == 'I' ||
		x == 'U' || x == 'E' || x == 'O'

	// Mengecek apakah karakter tersebut adalah huruf DAN termasuk golongan huruf vokal
	if huruf && (vKecil || vBesar) {
		// Menampilkan pesan "vokal" jika kondisi terpenuhi
		fmt.Println("vokal")
	// Mengecek apakah karakter tersebut adalah huruf TETAPI bukan golongan huruf vokal (artinya huruf konsonan)
	} else if huruf && !(vKecil || vBesar) {
		// Menampilkan pesan "konsonan"
		fmt.Println("konsonan")
	// Jalur alternatif jika karakter yang dimasukkan bukan merupakan huruf alfabet (misal: angka, spasi, atau simbol)
	} else {
		// Menampilkan pesan "bukan huruf"
		fmt.Println("bukan huruf")
	}
}