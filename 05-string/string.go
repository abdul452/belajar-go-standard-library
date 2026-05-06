package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Abdul Salim", "Abdul"))               // cek apakah string "Abdul" ada di dalam string "Abdul Salim"
	fmt.Println(strings.Split("Abdul Salim", " "))                      // memecah string "Abdul Salim" menjadi array dengan pemisah spasi
	fmt.Println(strings.ToLower("Abdul Salim"))                         // mengubah string "Abdul Salim" menjadi huruf kecil
	fmt.Println(strings.ToUpper("Abdul Salim"))                         // mengubah string "Abdul Salim" menjadi huruf besar
	fmt.Println(strings.Trim("   Abdul Salim   ", " "))                 // menghapus spasi di awal dan akhir string "   Abdul Salim   "
	fmt.Println(strings.ReplaceAll("Abdul Salim", "Salim", "Sulaiman")) // mengganti string "Salim" dengan "Sulaiman" dalam string "Abdul Salim"
}
