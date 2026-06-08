package readfile

import (
	"os"
	"strings"
)

func ReadFile(input, output string) (string, error) {
	data, err := os.ReadFile(input)
	if err != nil {
		return "error occured", err
	}
	word := string(data)

	text := strings.ToUpper(word)

	err = os.WriteFile(output, []byte(text), 0664)
	if err != nil {
		return "error occured", err
	}

	return word, nil
}
