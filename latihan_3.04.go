package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan nilai suhu dalam berbagai skala
	var celsius, fahrenheit, reamur, kelvin float64

	// Membaca input nilai suhu Celsius dari pengguna
	fmt.Scan(&celsius)

	// Konversi Celsius ke Fahrenheit: F = (C * 9/5) + 32
	fahrenheit = (celsius * 9 / 5) + 32

	// Konversi Celsius ke Reamur: R = C * 4/5
	reamur = celsius * 4 / 5

	// Konversi Celsius ke Kelvin: K = C + 273.15
	kelvin = celsius + 273.15

	// Menampilkan hasil konversi suhu tanpa angka di belakang koma (dibulatkan)
	fmt.Printf("Temperatur Celsius: %.0f\n", celsius)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}