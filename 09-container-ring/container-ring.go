package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := ring.New(5) // bikin ring dengan 5 elemen
	// cara isi data di ring secara manual
	// data.Value = "Aira"
	// data = data.Next() // pindah ke elemen berikutnya
	// data.Value = "Almahyra"
	// data = data.Next() // pindah ke elemen berikutnya
	// data.Value = "Falah"
	// data = data.Next() // pindah ke elemen berikutnya
	// data.Value = "Daffa"
	// data = data.Next() // pindah ke elemen berikutnya
	// data.Value = "Rizky"

	// cara isi data di ring dengan index manual
	// data.Value = "Aira"
	// data.Next().Value = "Almahyra"
	// data.Next().Next().Value = "Falah"
	// data.Next().Next().Next().Value = "Daffa"
	// data.Next().Next().Next().Next().Value = "Rizky"

	// cara isi data di ring dengan loop
	for i := 0; i < data.Len(); i++ {
		data.Value = fmt.Sprintf("Data ke-%d", i+1)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value.(string))
	})

	// kalau ambil data dengan index manual
	fmt.Println(data.Value.(string))                             // Aira
	fmt.Println(data.Next().Value.(string))                      // Almahyra
	fmt.Println(data.Next().Next().Value.(string))               // Falah
	fmt.Println(data.Next().Next().Next().Value.(string))        // Daffa
	fmt.Println(data.Next().Next().Next().Next().Value.(string)) // Rizky

}
