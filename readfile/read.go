package readfile

import (
	"os"
)

func ReadFile(input string) (string, error) {
	data, err := os.ReadFile(input)
	if err != nil {
		return "error occured", err
	}
	word := string(data)

	return word, nil
}
