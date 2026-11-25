//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel jam12 (format jam 12-jam), jam24 (format jam 24-jam) bertipe int, dan label bertipe string ("AM" atau "PM")
	var jam12, jam24 int
	var label string

	// Membaca input nilai jam dalam format 24-jam dari pengguna
	fmt.Scan(&jam24)

	// Percabangan menggunakan switch (switch expression-less) untuk mengonversi format 24-jam ke format 12-jam dan label waktu
	switch {
	// Kondisi 1: Jika jam menunjukkan pukul 00:00 (tengah malam)
	case jam24 == 0:
		// Dalam format 12-jam, pukul 0 dianggap sebagai pukul 12
		jam12 = 12
		// Memberikan label "AM" untuk waktu tengah malam / dini hari
		label = "AM"

	// Kondisi 2: Jika jam kurang dari 12 (pukul 01:00 hingga 11:59)
	case jam24 < 12:
		// Nilai jam 12-jam sama persis dengan nilai jam 24-jam
		jam12 = jam24
		// Memberikan label "AM" untuk waktu pagi hari
		label = "AM"

	// Kondisi 3: Jika jam tepat pukul 12:00 (siang hari)
	case jam24 == 12:
		// Dalam format 12-jam tetap bernilai 12
		jam12 = 12
		// Memberikan label "PM" untuk waktu siang hari
		label = "PM"

	// Kondisi 4: Jika jam lebih dari 12 (pukul 13:00 hingga 23:59)
	case jam24 > 12:
		// Mengurangi nilai jam 24-jam dengan 12 untuk mendapatkan konversi format 12-jam (misal: 13 menjadi 1)
		jam12 = jam24 - 12
		// Memberikan label "PM" untuk waktu sore hingga malam hari
		label = "PM"
	}

	// Menampilkan hasil akhir konversi jam format 12-jam beserta labelnya (AM/PM) ke konsol
	fmt.Println(jam12, label)
}