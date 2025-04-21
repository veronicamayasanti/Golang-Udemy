package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("Hello/worl.go"))
	fmt.Println(filepath.Base("Hello/worl.go"))
	fmt.Println(filepath.Ext("Hello/worl.go"))
	fmt.Println(filepath.IsAbs("Hello/worl.go"))
	fmt.Println(filepath.IsLocal("Hello/worl.go"))
	fmt.Println(filepath.Join(".hai", "perklankan", "aku.co"))
}
