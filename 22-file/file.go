package main

import (
	"bufio"
	"io"
	"os"
)

func creatNewFile(name string, message string) error {
	/*
		untuk parameter OpenFile(nama file string, flag, mode permission)
		untuk flag bisa di cek di package os
		lalu untuk mode permission bisa di cek di https://chmod-calculator.com/
	*/
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	// creatNewFile("sample.log", "this is sample log")

	// result, _ := readFile("sample.log")
	// fmt.Println(result)

	addToFile("sample.log", "\nthis is add message")
}
