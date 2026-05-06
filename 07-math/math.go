package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))  // untuk pembulatan ke atas
	fmt.Println(math.Floor(1.60)) // untuk pembulatan ke bawah
	fmt.Println(math.Round(1.40)) // untuk pembulatan ke terdekat
	fmt.Println(math.Max(10, 12)) // untuk mencari nilai terbesar
	fmt.Println(math.Min(10, 12)) // untuk mencari nilai terkecil
	fmt.Println(math.Sqrt(16))    // untuk mencari akar kuadrat
	fmt.Println(math.Pow(2, 3))   // untuk mencari pangkat
	fmt.Println(math.Abs(-10))    // untuk mencari nilai mutlak
	fmt.Println(math.Mod(10, 3))  // untuk mencari sisa bagi
}
