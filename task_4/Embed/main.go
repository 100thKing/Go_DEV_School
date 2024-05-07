package main

import (
	_ "embed"
	"fmt"
)

//go:embed text_file.txt
var message string

func main() {
	fmt.Println(message)
}
