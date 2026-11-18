//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel berat (total berat dalam gram), kg (bagian kilogram), dan gram (sisa berat dalam gram)
	var berat, kg, gram int
	// Deklarasi variabel biayaKg (biaya dari bagian kg), biayaGram (biaya dari sisa gram), dan total (total biaya akhir)
	var biayaKg, biayaGram, total int

	// Menampilkan teks/prompt ke konsol sebelum meminta input berat parsel
	fmt.Print("Berat parsel (gram): ")
	// Membaca input nilai berat total dalam gram dari pengguna
	fmt.Scan(&berat)

	// Menghitung jumlah kilogram dengan membagi total berat dengan 1000 secara pembagian bulat
	kg = berat / 1000
	// Menghitung sisa berat dalam gram menggunakan operasi sisa bagi (modulo) 1000
	gram = berat % 1000

	// Menghitung biaya untuk bagian kilogram dengan tarif tetap Rp 10.000 per kg
	biayaKg = kg * 10000

	// Mengecek apakah berat total sudah melebihi 10 kg
	if kg > 10 {
		// Jika berat di atas 10 kg, sisa gram tidak dikenakan biaya tambahan (gratis / 0)
		biayaGram = 0
	// Jika berat 10 kg atau kurang, mengecek apakah sisa gram bernilai 500 gram atau lebih
	} else if gram >= 500 {
		// Jika sisa gram >= 500, dikenakan tarif Rp 5 per gram
		biayaGram = gram * 5
	// Jalur alternatif jika sisa gram kurang dari 500 gram
	} else {
		// Jika sisa gram < 500, dikenakan tarif khusus yaitu Rp 15 per gram
		biayaGram = gram * 15
	}

	// Menghitung total keseluruhan biaya dengan menjumlahkan biaya bagian kilogram dan biaya sisa gram
	total = biayaKg + biayaGram

	// Menampilkan rincian hasil pemisahan berat (kg dan gram) ke konsol
	fmt.Println("Detail berat:", kg, "kg +", gram, "gr")
	// Menampilkan rincian nominal biaya berdasarkan komponen beratnya
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaGram)
	// Menampilkan hasil akhir total biaya pengiriman parsel
	fmt.Println("Total biaya: Rp.", total)
}