package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "A1, A2, A3\n" +
		"B1, B2, B3\n" +
		"C1, C2, C3\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for { // perulangan tanpa batas
		record, err := reader.Read() // balikan nya record per baris
		if err == io.EOF {           // end of file berarti data csv nya sudah habis
			break // selesai, keluar perulangan
		}

		fmt.Println(record)
	}
}
