# 📘 Repository Asprak-Algoritma-Program

Repositori ini berisi materi praktikum, contoh program, serta penyelesaian soal latihan untuk mata kuliah **Algoritma dan Pemrograman 1** menggunakan bahasa pemrograman **Go (Golang)** [cite: 1, 12]. Repositori ini disusun untuk mendukung kegiatan praktikum mahasiswa Program Studi S1 Informatika, Universitas Telkom [cite: 12, 13].

---

## 🏛️ Informasi Akademik & Institusi

* **Mata Kuliah** : Algoritma dan Pemrograman 1 [cite: 1, 12, 13]
* **Program Studi :** S1 Informatika [cite: 12, 13]
* **Fakultas :** Fakultas Informatika / School of Computing [cite: 12, 13]
* **Institusi :** Universitas Telkom [cite: 12, 13]
* **Penerbit Modul :** Informatics Lab, Universitas Telkom [cite: 12, 13]
* **Semester / Tahun Ajaran :** Semester Ganjil TA 2024/2025 [cite: 13]
* **Koordinator Mata Kuliah :** Prasti Eko Yunanto, S.T., M.Kom. (NIP. 19890017) [cite: 13]
* **Ketua Program Studi :** Dr. Erwin Budi Setiawan, S.Si., M.T. (NIP. 00760045) [cite: 13]

---

## 📚 Ringkasan Cakupan Modul

### 📖 Modul 2 : Input/Output (I/O), Tipe Data & Variabel
Modul 2 membahas konsep paling dasar dalam pemrograman Go, yaitu instruksi masukan (input), keluaran (output), penanganan tipe data dasar, manipulasi variabel, penggunaan operator, serta penyesuaian tipe data [cite: 1, 3, 4, 5].

#### Topik Utama :
* **📥📤 Input dan Output (I/O) :**
  * Pembacaan data dari pengguna menggunakan instruksi `fmt.Scan`, `fmt.Scanf`, dan `fmt.Scanln` [cite: 1, 7].
  * Menampilkan data ke layar monitor menggunakan `fmt.Print`, `fmt.Println`, dan `fmt.Printf` [cite: 2, 7].
* **🧠 Tipe Data & Variabel :**
  * Konsep variabel sebagai wadah penyimpanan data di lokasi memori [cite: 3].
  * Tipe data dasar Go: Integer (`int`, `int8`, `int32`, `int64`, `uint`, `uint8`, `uint32`, `uint64`), Real (`float32`, `float64`), Boolean (`bool`), Karakter (`byte`, `rune`), dan String (`string`) [cite: 3].
  * Penggunaan operator aritmatika, penggabungan string (`+`), bitwise (`&^`, `<<`, `>>`), komparasi, logika boolean (`&&`, `||`, `!`), serta operator pointer/memori (`&` dan `*`) [cite: 4].
  * Deklarasi variabel (`var a tipe` atau `a := nilai_awal`) serta pemahaman nilai bawaan (*default value*) untuk variabel yang tidak diinisialisasi [cite: 5, 6].
* **⚡ Konstanta Simbolik :**
  * Deklarasi konstanta bernilai tetap menggunakan kata kunci `const` [cite: 7].
* **🔤 Tipe Data Karakter :**
  * Penggunaan tipe data `byte` (uint8) dan `rune` (int32) berdasar acuan tabel ASCII/UTF-8 dan UTF-16 [cite: 7, 8].

#### 🎯 Contoh Soal & Pembahasan Modul 2 :
1. **Penjumlahan 5 Bilangan Bulat (`penjumlahan.go`) :** Membaca lima masukan bilangan bulat dan menampilkan hasil jumlah totalnya [cite: 8].
2. **Evaluasi Fungsi Matematika :** Menghitung nilai fungsi $f(x) = \frac{2}{x+5} + 5$ dengan tipe data riil [cite: 9].
3. **Konversi & Pergeseran ASCII (`ascii.go`) :** Membaca masukan integer untuk dicetak sebagai karakter ASCII serta melakukan pergeseran nilai karakter [cite: 9, 10].

#### 📂 Soal Latihan Modul 2 :
1. **Penelusuran Pertukaran Nilai String :** Mengamati dan menganalisis proses pertukaran nilai tiga variabel string (*swapping*) [cite: 10].
2. **Resume Biodata Mahasiswa :** Membaca nama, NIM, dan kelas mahasiswa lalu menampilkan format resume singkat [cite: 11].
3. **Perhitungan Luas Lingkaran :** Menghitung luas lingkaran berdasarkan masukan jari-jari $r$ bertipe riil [cite: 11].
4. **Konversi Suhu Fahrenheit ke Celcius :** Menghitung konversi suhu dari Fahrenheit ke Celcius dengan rumus $C = (F - 32) \times \frac{5}{9}$ [cite: 11].

---

### 📖 Modul 3 : Pendalaman I/O, Tipe Data & Variabel (Latihan 1)
Modul 3 merupakan pendalaman dari materi Modul 2 dengan fokus khusus pada *Integer Division*, operasi *Modulo*, serta teknik *Casting* atau konversi tipe data [cite: 12, 14, 15].

#### Topik Utama :
* **🔢 Integer Division (`div`) & Modulo (`mod`) :**
  * Operasi pembagian integer yang mengabaikan bagian pecahan (*floating point*) [cite: 14].
  * Operasi modulo (`%`) untuk memperoleh sisa pembagian integer [cite: 14].
  * Formula hubungan `dividend = quotient x divisor + remainder` serta penerapan pola `div` dan `mod` untuk isolasi digit bilangan [cite: 14, 15].
* **🔄 Casting & Konversi Tipe Data :**
  * Konversi eksplisit antar tipe data numerik (seperti `int(float_val)`) [cite: 15].
  * Penggunaan fungsi dari paket `strconv` (`strconv.Atoi` dan `strconv.Itoa`) untuk konversi antara `string` dan `int` [cite: 15].

#### 🎯 Contoh Soal & Pembahasan Modul  3 :
1. **Perhitungan Volume Kubus :** Menghitung volume kubus ($S^3$) berdasarkan masukan panjang sisi [cite: 16].
2. **Perhitungan Luas Segitiga :** Menghitung luas segitiga ($0.5 \times \text{alas} \times \text{tinggi}$) dari masukan alas dan tinggi [cite: 16, 17].
3. **Konversi Mata Uang (IDR ke USD) :** Menghitung konversi mata uang Rupiah ke Dolar AS dengan nilai kurs 15.000 IDR / USD [cite: 17, 18].

#### 📂 Soal Latihan Modul 3 :
1. **Pencarian Nilai $x$ dari Fungsi $f(x)$ :** Menentukan nilai $x$ apabila diketahui nilai $f(x)$ pada persamaan $f(x) = \frac{2}{x+5} + 5$ [cite: 19].
2. **Perhitungan Volume dan Luas Permukaan Bola :** Menghitung volume bola ($\frac{4}{3}\pi r^3$) dan luas kulit bola ($4\pi r^2$) dari masukan jari-jari $r$ [cite: 19].
3. **Pengecekan Tahun Kabisat :** Menentukan apakah suatu tahun merupakan tahun kabisat (habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100) dan menghasilkan nilai keluaran boolean (`true` / `false`) [cite: 19].
4. **Konversi Temperatur Celcius :** Menghitung konversi temperatur dari Celcius ke satuan Fahrenheit, Reamur, dan Kelvin [cite: 20].

---

## 🛠️ Cara Menjalankan Program Go

Pastikan lingkungan pemrogranan Go telah terpasang pada komputer Anda[cite: 2, 8].

1. **Menjalankan Kode Program Langsung :**
   ```bash
   go run nama_file.go
