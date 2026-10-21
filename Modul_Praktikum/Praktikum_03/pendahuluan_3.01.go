//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi konstanta bobot persentase nilai untuk setiap komponen penilaian
	const (
		tugas     float32 = 0.15 // Bobot Tugas = 15%
		kuis      float32 = 0.15 // Bobot Kuis = 15%
		praktikum float32 = 0.20 // Bobot Praktikum = 20%
		uts       float32 = 0.25 // Bobot UTS = 25%
		uas       float32 = 0.25 // Bobot UAS = 25%
	)

	// Definisi struct mahasiswa untuk mengelompokkan data identitas dan komponen nilai
	type mahasiswa struct {
		nama       string
		nim        string
		nTugas     float32
		nKuis      float32
		nPraktikum float32
		nUts       float32
		nUas       float32
		nAkhir     float32
	}

	// Deklarasi variabel DataMahasiswa bertipe struct mahasiswa
	var DataMahasiswa mahasiswa

	// Membaca input data mahasiswa (nama, NIM, serta komponen nilai satu per satu)
	fmt.Scan(&DataMahasiswa.nama, &DataMahasiswa.nim, &DataMahasiswa.nTugas, &DataMahasiswa.nKuis,
		&DataMahasiswa.nPraktikum, &DataMahasiswa.nUts, &DataMahasiswa.nUas)

	// Menghitung akumulasi nilai akhir berdasarkan penjumlahan tiap komponen nilai dikali bobotnya
	DataMahasiswa.nAkhir = DataMahasiswa.nTugas*tugas + DataMahasiswa.nKuis*kuis +
		DataMahasiswa.nPraktikum*praktikum + DataMahasiswa.nUts*uts +
		DataMahasiswa.nUas*uas

	// Menampilkan output ringkasan data mahasiswa dan detail komponen nilainya (2 angka di belakang koma)
	fmt.Println("--- Data Mahasiswa ---")
	fmt.Printf("NIM : %s\n", DataMahasiswa.nim)
	fmt.Printf("Nama : %s\n", DataMahasiswa.nama)
	fmt.Printf("Nilai Tugas : %.2f\n", DataMahasiswa.nTugas)
	fmt.Printf("Nilai Kuis : %.2f\n", DataMahasiswa.nKuis)
	fmt.Printf("Nilai Praktikum : %.2f\n", DataMahasiswa.nPraktikum)
	fmt.Printf("Nilai UTS : %.2f\n", DataMahasiswa.nUts)
	fmt.Printf("Nilai UAS : %.2f\n", DataMahasiswa.nUas)
	fmt.Printf("Nilai Akhir : %.2f\n", DataMahasiswa.nAkhir)
}