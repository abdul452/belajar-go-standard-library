package main

import "container/list"

func main() {
	data := list.New()
	data.PushBack("Aira")
	data.PushBack("Almahyra")
	data.PushBack("Falah")

	for element := data.Front(); element != nil; element = element.Next() {
		println(element.Value.(string))
	}

	// kalau ambil data dengan index manual
	head := data.Front()
	println(head.Value.(string)) // Aira

	second := head.Next()
	println(second.Value.(string)) // Almahyra

	third := second.Next()
	println(third.Value.(string)) // Falah
}
