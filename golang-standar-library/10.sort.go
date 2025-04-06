package main

import (
	"fmt"
	"sort"
)

//	package sort berisikan utilitas untuk proses pengurutan
//	agar data bisa diurutkan harus mengimplementasikan kontrak di interface sort.interface

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]

}

func main() {
	users := []User{
		{"Bob", 21},
		{"Alice", 20},
		{"rani", 27},
		{"lady", 24},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
