package main

import (
	"fmt"
	"hackathon/readfile"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Argument can't be less than 2")
		return
	}
	str := os.Args[1]
	output := os.Args[2]

	_, err := readfile.ReadFile(str, output)
	if err != nil {
		fmt.Printf("occured: %v", err)
	}
	fmt.Println("ALl done")
}
