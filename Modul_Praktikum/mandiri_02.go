package main

import "fmt"

func main() {
	// Deklarasi variabel untuk komponen gaji dan jam lembur bertipe int
	var gaji_pokok, potongan, bonus_lembur, jam_lembur, gaji_bersih int

	// Membaca input nilai gaji pokok dan total jam lembur dari pengguna
	fmt.Scan(&gaji_pokok, &jam_lembur)

	// Menghitung total bonus lembur (Rp45.000 per jam)
	bonus_lembur = 45000 * jam_lembur

	// Menghitung total potongan gaji (potongan 2% + potongan 3.5%)
	potongan = ((2 * gaji_pokok) / 100) + ((35 * gaji_pokok) / 1000)

	// Menghitung total gaji bersih setelah ditambah bonus dan dikurangi potongan
	gaji_bersih = gaji_pokok + bonus_lembur - potongan

	// Menampilkan hasil akhir perhitungan gaji bersih
	fmt.Println(gaji_bersih)
}