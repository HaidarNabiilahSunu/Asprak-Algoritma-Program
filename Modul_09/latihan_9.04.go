//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel aksi (pilihan moda transportasi) bertipe string
	var aksi string
	// Deklarasi variabel status kondisi bajuBasah dan telat bertipe bool (nilai awal false)
	var bajuBasah, telat bool

	// Membaca input jenis aksi/transportasi dari pengguna (misal: "motor" atau "mobil")
	fmt.Scan(&aksi)

	// Inisialisasi awal status kondisi menjadi false sebelum dilakukan pengecekan
	bajuBasah = false
	telat = false

	// Mengecek jika pilihan aksi adalah menggunakan "motor"
	if aksi == "motor" {
		// Mengubah status bajuBasah menjadi true karena naik motor berisiko basah (terkena hujan)
		bajuBasah = true
		// Menampilkan pesan spesifik terkait kondisi menggunakan motor
		fmt.Println("Baju basah")
	}

	// Mengecek jika pilihan aksi adalah menggunakan "mobil"
	if aksi == "mobil" {
		// Mengubah status telat menjadi true karena naik mobil berisiko terkena kemacetan
		telat = true
		// Menampilkan pesan spesifik terkait risiko terlambat
		fmt.Println("Terlambat masuk kuliah")
	}

	// Mengecek jika status bajuBasah bernilai true
	if bajuBasah {
		// Menampilkan konsekuensi tambahan yang harus dilakukan jika baju basah
		fmt.Println("Konsekuensi : Harus membawa baju salin")
	}

	// Mengecek jika status telat bernilai true
	if telat {
		// Menampilkan konsekuensi akademik yang didapat jika terlambat
		fmt.Println("Konsekuensi : Mendapatkan pengurangan nilai")
	}
}