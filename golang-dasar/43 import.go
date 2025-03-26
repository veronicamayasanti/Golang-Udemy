package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	result := helper.SayHello("maya")
	fmt.Println(result)

	fmt.Println(helper.Application)
	//	paksa jalankan maka akan undefined
	//fmt.Println(helper.version)
	//fmt.Println(helper.sayGoodBye("santi"))
}
