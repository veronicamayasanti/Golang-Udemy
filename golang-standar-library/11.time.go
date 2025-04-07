package main

import (
	"fmt"
	"time"
)

// package time berisikan fungsional untuk management waktu di go-lang
//	beberapa function di package time
//	function				kegunaan
//	time.Now()				untuk mendapat waktu saat ini
//	time.Date(...)			untuk membuat waktu
//	time.Parse(layout, string)		untuk memparsing waktu dari string

func main() {
	now := time.Now()
	fmt.Println("waktu lokal : ", now.Local())

	utc := time.Date(2025, time.April, 7, 22, 4, 0, 0, time.UTC)
	fmt.Println("waktu lokal utc : ", utc.Local())

	//formatter := "2006-01-02 15:16:05"
	//value := "2020-10-01 10:15:10"
	//value := "salah"

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("error : ", err.Error())
	} else {
		fmt.Println(valueTime)
	}
	fmt.Println(valueTime.Year())

}
