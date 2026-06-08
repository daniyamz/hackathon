package main

import (
	"fmt"
	"strconv"
)

func BinToDec(str string) (int64, error) {
	dec, _ := strconv.ParseInt(str, 2, 64)
	return dec, nil
}

func main() {
	binres, err := BinToDec("11111111")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(binres)
}
