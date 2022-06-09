package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var unpacked strings.Builder

	if input == "" {
		return "", nil
	}

	fmt.Println("input", input)

	for index, character := range input {
		if index == 0 {
			unpacked.WriteString(string(character))
		} else {
			digit, _ := strconv.Atoi(string(character))
			previous := string(input[index-1])

			if unicode.IsDigit(character) && digit > 0 {
				letters := strings.Repeat(previous, digit-1)
				unpacked.WriteString(letters)
			} else {
				unpacked.WriteString(string(character))
			}
		}
		fmt.Println(string(character), unpacked.String())

	}

	return unpacked.String(), nil
}
