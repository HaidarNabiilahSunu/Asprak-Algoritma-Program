package main

import "fmt"

func main() {
	// Deklarasi variabel fahrenheit dan celsius dengan tipe data float64
	var fahrenheit, celsius float64

	// Membaca input nilai suhu dalam derajat Fahrenheit dari pengguna
	fmt.Scan(&fahrenheit)

	// Konversi suhu dari Fahrenheit ke Celsius menggunakan rumus C = (F - 32) * 5 / 9
	celsius = (fahrenheit - 32) * 5 / 9

	// Menampilkan hasil konversi ke Celsius tanpa angka di belakang koma (dibulatkan)
	fmt.Printf("%.0f\n", celsius)
}