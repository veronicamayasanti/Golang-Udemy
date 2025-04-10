package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "maya,santi,tata\n" +
		"budi,pratama,geonathan\n" +
		"alex,rasmus,alya"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {

			break
		}
		fmt.Println(record)
	}

}
