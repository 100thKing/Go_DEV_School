package main

//go:generate go build -tags=gen
//go:generate ./Go_Generate

import (
	"fmt"
)

var message string = "Runned without generation"

func main() {
	fmt.Println(message)
}
