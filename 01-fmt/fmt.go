package main

import "fmt"

func main() {
	firstName := "Abdul"
	lastName := "Salim"
	age := 30

	// menggunakan fmt.Printf untuk mencetak string yang diformat
	fmt.Printf("My name is %s %s and I am %d years old.\n", firstName, lastName, age)

	// menggunakan fmt.Sprintf untuk membuat string yang diformat dan menyimpannya dalam variabel
	formattedString := fmt.Sprintf("My name is %s %s and I am %d years old.", firstName, lastName, age)
	fmt.Println(formattedString)
}
