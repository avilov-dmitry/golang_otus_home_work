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

	inputUTF := []rune(input)

	for index, character := range inputUTF {
		lastIndex := len(inputUTF) - 2

		if index > lastIndex {
			if !unicode.IsDigit(character) {
				unpacked.WriteString(string(character))
			}
			break
		}

		nextCharacter := inputUTF[index+1]
		isDigit := unicode.IsDigit(character)
		isDigitNext := unicode.IsDigit(nextCharacter)

		if (isDigit && index == 0) || (isDigit && isDigitNext) {
			return "", ErrInvalidString
		}

		if isDigit {
			continue
		}

		if !isDigitNext {
			unpacked.WriteRune(character)
			continue
		}

		digit, err := strconv.Atoi(string(nextCharacter))
		if err != nil {
			return "", ErrInvalidString
		}

		if digit > 0 {
			for i := 0; i < digit; i++ {
				unpacked.WriteRune(character)
			}
		}
	}

	return unpacked.String(), nil
}
