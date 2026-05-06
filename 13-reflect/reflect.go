package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"15` // ini pakai fitur struct tag misal untuk validasi
}

type Person struct {
	Name    string `required:"true" max:"10`
	Address string `required:"true" max:"10`
	Email   string `required:"true" max:"10`
}

func IsValid(value any) (result bool) { // contoh validation menggunakan struct tag
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ { // cek 1 1 field nya
		f := t.Field(i)
		if f.Tag.Get("required") == "true" { // jika field ersebut required
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""  // cek datanya kosong apa engga
			if result == false { // jika kosong
				return result // maka langsung return false
			}
		}
	}
	return true
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println()
	}
}

func main() {
	readField(Sample{"Abdul salim salim salim"})
	readField(Person{"Abdul", "", ""})

	person := Person{
		Name:    "Abdul",
		Address: "",
		Email:   "asd",
	}
	fmt.Println(IsValid(person))
}
