//go:build ignore
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan data nama, NIM, dan kelas
	var nama, nim, kelas string

	// Membaca input secara berurutan untuk variabel nama, nim, dan kelas
	fmt.Scan(&nama, &nim, &kelas)

	// Menampilkan kalimat perkenalan dengan menggabungkan nilai dari variabel
	fmt.Println("Perkenalkan saya adalah", nama,
		", salah satu mahasiswa Prodi S1-IF dari kelas",
		kelas, "dengan NIM", nim+".")
}