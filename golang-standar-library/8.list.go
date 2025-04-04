package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("santi")
	data.PushBack("jack")
	data.PushBack("mike")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // santi

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	//	perulangan
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
