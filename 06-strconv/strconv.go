package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true") // konversi string "true" menjadi boolean true
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Result", result)
	}

	resultInt, err := strconv.Atoi("1000") // konversi string "1000" menjadi integer 1000
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Result", resultInt)
	}

	binary := strconv.FormatInt(1000, 2) // konversi integer 1000 menjadi string biner "1111101000"
	fmt.Println("Binary", binary)

	octal := strconv.FormatInt(1000, 8) // konversi integer 1000 menjadi string oktal "1750"
	fmt.Println("Octal", octal)

	hexadecimal := strconv.FormatInt(1000, 16) // konversi integer 1000 menjadi string heksadesimal "3e8"
	fmt.Println("Hexadecimal", hexadecimal)

	stringInt := strconv.Itoa(1000) // konversi integer 1000 menjadi string "1000"
	fmt.Println("String Integer", stringInt)

}
