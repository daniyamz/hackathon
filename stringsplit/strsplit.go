package main

import (
	"fmt"
	"strings"
	"unicode"
)

func UseFields(input string) ([]string, error) {
	fieldstr := strings.Fields(input)
	return fieldstr, nil
}

func StrSplit(s string) ([]string, error) {
	var word strings.Builder
	text := []string{}
	ssplit := strings.Split(s, " ")
	for _, val := range ssplit {
		for _, char := range val {
			if unicode.IsPunct(char) {
				if word.Len() != 0 {
					text = append(text, word.String())
					word.Reset()
				}
				text = append(text, string(char))
				continue
			} else if char != ' ' {
				word.WriteString(string(char))
			}
		}
	}
	return text, nil
}
func main() {
	d, err := UseFields("Hello, world! How are you?")
	if err != nil {
		fmt.Println("Erorr", err)
	}
	r, err := StrSplit("Hello, world! How are you?")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(d)
	fmt.Println(r)
}
