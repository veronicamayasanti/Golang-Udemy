package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"maya", "santi", "veronica"})
	_ = writer.Write([]string{"tata", "alya", "michel"})
	_ = writer.Write([]string{"geo", "nathan", "wijaya"})

	writer.Flush()
}
