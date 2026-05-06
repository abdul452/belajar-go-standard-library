package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// contract dari package sort
func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{"Abdul", 30},
		{"Dinda", 29},
		{"Aira", 2},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
