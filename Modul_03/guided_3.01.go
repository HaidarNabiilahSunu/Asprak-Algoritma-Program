package main

import "fmt"

func main() {
	var sisi, volume float64

	fmt.Printf("Masukan Sisi : ")
	fmt.Scan(&sisi)

	volume = sisi * sisi * sisi
	fmt.Printf("Jawabanya = %.2f\n", volume)
}