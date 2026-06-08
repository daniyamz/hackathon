package main

import (
	"fmt"
	"hackathon/readfile"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Argument can't be less than 2")
		return
	}
	str := os.Args[1]

	cont, err := readfile.ReadFile(str)
	if err != nil {
		fmt.Printf("occured: %v", err)
	}
	fmt.Println(cont)
}
