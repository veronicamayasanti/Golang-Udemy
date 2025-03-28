package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFound        = errors.New("resource not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "maya" {
		return NotFound
	}
	//	sukses
	return nil
}

func main() {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFound) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unkown error")
		}
	}
}
