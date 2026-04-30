package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("data not found")
)

func GetByName(name string) error {
	if name == "" {
		return ValidationError
	}

	if name != "John" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetByName("")
	if err != nil {
		// pakai errors.Is untuk memeriksa jenis error
		if errors.Is(err, ValidationError) { // memeriksa apakah error adalah ValidationError
			fmt.Println("Terjadi kesalahan validasi:", err)
		} else if errors.Is(err, NotFoundError) { // memeriksa apakah error adalah NotFoundError
			fmt.Println("Data tidak ditemukan:", err)
		} else {
			fmt.Println("Kesalahan lainnya:", err)
		}
	} else {
		fmt.Println("Data ditemukan dengan nama John")
	}
}
