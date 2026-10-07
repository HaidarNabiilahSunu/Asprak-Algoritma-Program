//go:build ignore
package main

import "fmt"

// Tipe bentukan untuk menyimpan data investasi saham
type Investasi struct {
	hargaBeli       float64
	hargaJual       float64
	jumlahSaham     float64
	totalInvestasi  float64
	totalPenjualan  float64
	keuntunganKotor float64
	biayaTransaksi  float64
	pajak           float64
	keuntunganBersih float64
}

func main() {
	var i Investasi

	// Input harga beli, harga jual, dan jumlah saham
	fmt.Print("Harga Beli: ")
	fmt.Scan(&i.hargaBeli)

	fmt.Print("Harga Jual: ")
	fmt.Scan(&i.hargaJual)

	fmt.Print("Jumlah Saham: ")
	fmt.Scan(&i.jumlahSaham)

	// Menghitung total investasi awal
	i.totalInvestasi = i.hargaBeli * i.jumlahSaham

	// Menghitung total penjualan
	i.totalPenjualan = i.hargaJual * i.jumlahSaham

	// Menghitung keuntungan kotor
	i.keuntunganKotor = i.totalPenjualan - i.totalInvestasi

	// Biaya transaksi sebesar 0,2% dari total penjualan
	i.biayaTransaksi = 0.002 * i.totalPenjualan

	// Pajak keuntungan sebesar 10%
	// Jika keuntungan negatif, pajak menjadi 0
	if i.keuntunganKotor > 0 {
		i.pajak = 0.10 * i.keuntunganKotor
	} else {
		i.pajak = 0
	}

	// Menghitung keuntungan bersih
	i.keuntunganBersih = i.keuntunganKotor - i.biayaTransaksi - i.pajak

	// Menampilkan hasil perhitungan
	fmt.Printf("Total Investasi Awal: Rp %.2f\n", i.totalInvestasi)
	fmt.Printf("Total Penjualan: Rp %.2f\n", i.totalPenjualan)
	fmt.Printf("Keuntungan Kotor: Rp %.2f\n", i.keuntunganKotor)
	fmt.Printf("Biaya Transaksi: Rp %.2f\n", i.biayaTransaksi)
	fmt.Printf("Pajak Keuntungan: Rp %.2f\n", i.pajak)
	fmt.Printf("Keuntungan Bersih: Rp %.2f\n", i.keuntunganBersih)
}