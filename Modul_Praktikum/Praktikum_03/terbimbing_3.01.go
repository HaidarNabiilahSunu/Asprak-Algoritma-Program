//go:build ignore
package main

import "fmt"

func main() {
	// Definisi struct Book untuk mengelompokkan atribut/informasi dari sebuah buku
	type Book struct {
		title       string
		genre       string
		releaseYear int
		rating      float64
	}

	// Deklarasi dua variabel (book1 dan book2) bertipe struct Book
	var book1, book2 Book

	// Membaca input judul untuk buku pertama dan buku kedua
	fmt.Scan(&book1.title, &book2.title)

	// Membaca input genre untuk buku pertama dan buku kedua
	fmt.Scan(&book1.genre, &book2.genre)

	// Membaca input tahun rilis untuk buku pertama dan buku kedua
	fmt.Scan(&book1.releaseYear, &book2.releaseYear)

	// Membaca input rating untuk buku pertama dan buku kedua
	fmt.Scan(&book1.rating, &book2.rating)

	// Menampilkan data dan informasi detail dari buku pertama
	fmt.Println("--- Book 1 ---")
	fmt.Println("Book Title :", book1.title)
	fmt.Println("Book Genre :", book1.genre)
	fmt.Println("Release Year :", book1.releaseYear)
	fmt.Printf("Rating : %.2f\n", book1.rating)

	// Menampilkan data dan informasi detail dari buku kedua
	fmt.Println("--- Book 2 ---")
	fmt.Println("Book Title :", book2.title)
	fmt.Println("Book Genre :", book2.genre)
	fmt.Println("Release Year :", book2.releaseYear)
	fmt.Printf("Rating : %.2f\n", book2.rating)
}