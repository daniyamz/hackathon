package main

import (
	"fmt"
	"strconv"
)

func HexConvt(input string) (int64, error) {
	digit, _ := strconv.ParseInt(input, 16, 64)

	return digit, nil
}

func main() {
	res, err := HexConvt("BADF00D")
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(res)
	fmt.Println(binres)
}
