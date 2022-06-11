package hw02unpackstring

import (
	"errors"
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

	for index, character := range input {
		if index > len(input)-2 {
			if !unicode.IsDigit(character) {
				unpacked.WriteString(string(character))
			}
			break
		}

		nextCharacter := input[index+1]
		digit, _ := strconv.Atoi(string(nextCharacter))
		isDigit := unicode.IsDigit(character)
		isDigitNext := unicode.IsDigit(rune(nextCharacter))

		if (unicode.IsDigit(character) && index == 0) || (isDigit && isDigitNext) {
			return "", ErrInvalidString
		}

		if isDigitNext {
			if digit > 0 {
				letters := strings.Repeat(string(character), digit)
				unpacked.WriteString(letters)
			}
		} else {
			if !unicode.IsDigit(character) {
				unpacked.WriteString(string(character))
			}
		}
	}

	return unpacked.String(), nil
}
