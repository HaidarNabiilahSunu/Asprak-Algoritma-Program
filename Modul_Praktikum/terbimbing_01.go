package main

import "fmt"

func main() {
	// Deklarasi variabel untuk kecepatan (km/jam), durasi menit, durasi jam, dan total jarak (km)
	var kecepatan, menit int
	var jam, total_jarak int

	// Membaca input nilai kecepatan dari pengguna
	fmt.Scan(&kecepatan)

	// Menghitung total jarak keseluruhan dari penjumlahan beberapa segmen perjalanan (100 + 60 + 170 km)
	total_jarak = 100 + 60 + 170

	// Menghitung total waktu tempuh dalam menit (Jarak / Kecepatan * 60)
	menit = (total_jarak * 60) / kecepatan

	// Konversi total menit menjadi jam
	jam = menit / 60

	// Menghitung sisa menit setelah dikurangi nilai jam
	menit = menit % 60

	// Menampilkan hasil durasi perjalanan dalam format jam dan menit
	fmt.Println(jam, menit)
}