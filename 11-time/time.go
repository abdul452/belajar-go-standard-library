package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // untuk mendapatkan waktu saat ini
	fmt.Println(now.Local())

	utc := time.Date(2026, time.May, 15, 14, 30, 23, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// time.RFC3339
	formatter := "2006-01-01 15:01:03" // cek di doc RFC3339 (ctrl + klik)8

	value := "2022-10-10 10:10:10" // misal ini inputan dari user
	// value := "ASAL"
	valuetime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valuetime)
		fmt.Println(valuetime.Year())
		fmt.Println(valuetime.Month())
		fmt.Println(valuetime.Day())
		fmt.Println(valuetime.Hour())
	}
}
