package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"A1", "A2", "A3"})
	_ = writer.Write([]string{"B1", "B2", "B3"})
	_ = writer.Write([]string{"C1", "C2", "C3"})

	writer.Flush() // mengirim semua perubahan
}
