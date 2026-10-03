package main

import "fmt"

func main() {
	var alas, tinggi int
	var luas float64

	fmt.Printf("Masukan Alas : ")
	fmt.Scan(&alas)

	fmt.Printf("Masukan Tinggi : ")
	fmt.Scan(&tinggi)

	luas = 0.5 * float64(alas*tinggi)

	fmt.Println("Jawabanya = %.2f\n",luas)
}