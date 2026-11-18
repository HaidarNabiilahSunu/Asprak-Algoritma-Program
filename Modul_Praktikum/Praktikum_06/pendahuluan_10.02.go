//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel t (total pertemuan), h (jumlah kehadiran), dan k (jumlah tugas) bertipe int
	var t, h, k int
	// Deklarasi variabel hasil bertipe string untuk menyimpan keterangan status izin ujian
	var hasil string

	// Membaca input nilai t, h, dan k secara berurutan dari pengguna
	fmt.Scan(&t, &h, &k)

	// Menetapkan nilai awal default hasil menjadi "TIDAK BOLEH UAS"
	hasil = "TIDAK BOLEH UAS"

	// Mengecek apakah persentase kehadiran minimal 75% (h*100 >= 75*t) DAN jumlah tugas (k) minimal 8
	if h*100 >= 75*t && k >= 8 {
		// Mengubah status hasil menjadi "BOLEH UAS" jika kedua syarat terpenuhi
		hasil = "BOLEH UAS"
	}

	// Menampilkan hasil akhir status kelayakan mengikuti UAS ke konsol
	fmt.Println(hasil)
}