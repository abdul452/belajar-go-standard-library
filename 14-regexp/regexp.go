package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Pola harus diawali 'a', diakhiri 'l', dan tengahnya huruf kecil semua
	var regex *regexp.Regexp = regexp.MustCompile(`^a[a-z]+l$`)

	// return nya boolean
	fmt.Println(regex.MatchString("adl"))
	fmt.Println(regex.MatchString("abdul"))
	fmt.Println(regex.MatchString("aBDUL"))

	/* return nya slice
	ini ga dapet karna
	Tanda ^ memaksa regex untuk mencari kecocokan mulai dari karakter pertama di string,
	dan $ memaksa kecocokan berakhir di karakter terakhir
	Git mendeteksi string Anda sebagai satu kesatuan utuh.
	Karena di tengah-tengah string tersebut ada spasi dan kata-kata lain,
	maka keseluruhan string "abdul adl Abdul" tidak memenuhi syarat "diawali a
	// dan langsung diakhiri l */
	fmt.Println(regex.FindAllString("abdul adl Abdul", 10)) // dari string yang dicari ingin maksimal dapetin slice nya 10

	// kalau ini baru bisa
	regex = regexp.MustCompile(`a[a-z]+l`)
	fmt.Println(regex.FindAllString("abdul adl Abdul", 10)) // dari string yang dicari ingin maksimal dapetin slice nya 10

	// Contoh string HTML hasil crawl
	htmlContent := `<div>Konten luar</div><div2>Data Penting Anda</div2><div3>Footer</div3>`

	// Regex:
	// <div2>   : Mencari teks awal
	// (.*?)    : Menangkap isi di tengahnya (grup 1)
	// <div3>   : Mencari teks akhir
	regex = regexp.MustCompile(`<div2>(.*?)</div2>`)

	// 1. Mengambil data (Return-nya adalah slice)
	// FindStringSubmatch akan mengembalikan [full_match, group_1]
	result := regex.FindStringSubmatch(htmlContent)

	if len(result) > 1 {
		fmt.Println("Data yang ditemukan:", result[1])
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
