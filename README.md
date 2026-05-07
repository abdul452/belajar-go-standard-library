# Belajar Go Standard Library

Repositori ini berisi kumpulan kode latihan dan implementasi berbagai paket dalam **Standard Library Go**. Fokus utama dari repositori ini adalah untuk memahami fungsi-fungsi bawaan (*built-in*) yang disediakan oleh Go guna membangun aplikasi yang efisien dan tangguh tanpa ketergantungan berlebihan pada library pihak ketiga.

## 🎯 Objektif Pembelajaran
- Memahami manipulasi data dasar (strings, math, sort, container).
- Mengelola input/output dan sistem file (os, path, bufio, file).
- Mengolah format data populer (csv, base64).
- Mempelajari teknik lanjutan seperti Reflection dan Regular Expressions.

## 📁 Struktur Paket & Materi
Repositori ini disusun secara modular berdasarkan paket yang ada di Go Standard Library:

### Dasar & Utiliti
- `01-fmt`: Formatting I/O (Print, Scan, dll).
- `02-errors`: Penanganan error dan pembuatan custom error.
- `05-string` & `06-strconv`: Manipulasi string dan konversi tipe data.
- `07-math`: Fungsi matematika dasar.
- `10-sort`: Pengurutan data pada slice dan array.

### Sistem & File
- `03-os`: Interaksi dengan sistem operasi dan environment variables.
- `04-flag`: Pengelolaan argumen command-line.
- `19-path`: Manipulasi path file.
- `20-bufio-read` & `21-bufio-writer`: Buffered I/O untuk performa baca/tulis yang lebih baik.
- `22-file`: Operasi file secara langsung.

### Struktur Data Lanjutan
- `08-container-list`: Implementasi Doubly Linked List.
- `09-container-ring`: Implementasi Circular List.
- `18-slices`: Utiliti modern untuk manipulasi slice (Go terbaru).

### Data & Waktu
- `11-time` & `12-duration`: Pengelolaan waktu, tanggal, dan durasi.
- `14-regexp`: Pemrosesan teks menggunakan Regular Expressions.
- `15-encod-base64`: Encoding dan decoding data ke format Base64.
- `16-csv-reader` & `17-csv-writer`: Membaca dan menulis file format CSV.

### Konsep Lanjutan
- `13-reflect`: Penggunaan Reflection untuk inspeksi tipe data saat runtime.

## ⚙️ Cara Menjalankan
1. Pastikan Go sudah terinstal di sistem Anda (disarankan versi 1.21 ke atas).
2. Clone repositori ini:
   ```bash
   git clone [https://github.com/abdul452/belajar-go-standard-library.git](https://github.com/abdul452/belajar-go-standard-library.git)
3. Masuk ke folder modul yang diinginkan dan jalankan:
  ```bash
    cd 14-regexp
    go run main.go
