//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel nama_tanaman bertipe string untuk menampung nama tanaman masukan dari pengguna
	var nama_tanaman string

	// Membaca input string berupa nama tanaman dari pengguna
	fmt.Scan(&nama_tanaman)

	// Percabangan switch-case untuk mengklasifikasikan jenis tanaman berdasarkan nama yang dimasukkan
	switch nama_tanaman {
	// Kasus 1: Jika tanaman yang dimasukkan adalah "nepenthes" atau "drosera"
	case "nepenthes", "drosera":
		// Menampilkan informasi bahwa tanaman tersebut merupakan bagian dari tanaman karnivora
		fmt.Println("Termasuk Tanaman Karnivora.")
		// Menampilkan informasi asal tanaman yang merupakan flora asli Indonesia
		fmt.Println("Asli Indonesia.")

	// Kasus 2: Jika tanaman yang dimasukkan adalah "venus" atau "sarracenia"
	case "venus", "sarracenia":
		// Menampilkan informasi bahwa tanaman tersebut merupakan bagian dari tanaman karnivora
		fmt.Println("Termasuk Tanaman Karnivora.")
		// Menampilkan informasi bahwa tanaman tersebut bukan berasal dari Indonesia
		fmt.Println("Tidak Asli Indonesia.")

	// Kondisi default: Dijalankan jika nama tanaman tidak cocok dengan salah satu case di atas
	default:
		// Menampilkan pesan bahwa tanaman tersebut tidak dikategorikan sebagai tanaman karnivora
		fmt.Println("Tidak termasuk Tanaman Karnivora.")
	}
}