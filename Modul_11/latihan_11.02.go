//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel kendaraan (jenis kendaraan), serta durasi (lama parkir), tarif (harga per jam), dan total (total biaya keseluruhan)
	var kendaraan string
	var durasi int
	var tarif, total int

	// Membaca input jenis kendaraan dan durasi parkir dari pengguna secara berurutan
	fmt.Scan(&kendaraan, &durasi)

	// Validasi batas minimum durasi: jika durasi kurang dari 1 jam, bulatkan atau tetapkan minimal menjadi 1 jam
	if durasi < 1 {
		durasi = 1
	}

	// Percabangan switch untuk menentukan besaran tarif dasar per jam berdasarkan jenis kendaraan
	switch kendaraan {
	// Jika jenis kendaraan adalah "motor", tarif dasarnya adalah Rp 2.000 per jam
	case "motor":
		tarif = 2000

	// Jika jenis kendaraan adalah "mobil", tarif dasarnya adalah Rp 5.000 per jam
	case "mobil":
		tarif = 5000

	// Jika jenis kendaraan adalah "truk", tarif dasarnya adalah Rp 8.000 per jam
	case "truk":
		tarif = 8000

	// Blok default: dijalankan apabila jenis kendaraan tidak dikenali/tidak valid, tarif diset 0
	default:
		tarif = 0
	}

	// Menghitung total biaya parkir dengan mengalikan tarif per jam dengan total durasi
	total = tarif * durasi

	// Menampilkan hasil akhir jumlah total biaya parkir menggunakan format string (fmt.Printf)
	fmt.Printf("Rp %d\n", total)
}