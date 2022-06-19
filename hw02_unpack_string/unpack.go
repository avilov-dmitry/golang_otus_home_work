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
		isDigit := unicode.IsDigit(character)
		isDigitNext := unicode.IsDigit(rune(nextCharacter))

		if (isDigit && index == 0) || (isDigit && isDigitNext) {
			return "", ErrInvalidString
		}

		if isDigit {
			continue
		} else {
			if isDigitNext {
				digit, err := strconv.Atoi(string(nextCharacter))

				if err != nil {
					return "", ErrInvalidString
				}

				if digit > 0 {
					letters := strings.Repeat(string(character), digit)
					unpacked.WriteString(letters)
				}
			} else {
				if !isDigit {
					unpacked.WriteString(string(character))
				}
			}
		}

	}

	return unpacked.String(), nil
}
