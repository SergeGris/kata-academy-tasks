
package main;

import (
	"errors"
	"bufio"
	"strings"
	"os"
	"fmt"
	"strconv"
)

func isRoman(c byte) bool {
	switch (c) {
	case 'I', 'V', 'X', 'L', 'C', 'D', 'M':
		return true
	default:
		return false
	}
}

var (
	ErrInvalidRomanDigit = errors.New("Invalid roman digit")
)

func fromRoman(roman string) (int, error) {
	translateRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var value, prevDigit int

	for i := len(roman) - 1; i >= 0; i-- {
		romanDigit := roman[i]

		if (!isRoman(romanDigit)) {
			return 0, ErrInvalidRomanDigit
		}

		decDigit := translateRoman[romanDigit]

		if decDigit < prevDigit {
			value -= decDigit
		} else {
			value += decDigit
			prevDigit = decDigit
		}
	}

	return value, nil
}

var (
	ErrTooBigValue = errors.New("Too big value")
)

func toRoman(number int) (string, error) {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return "", ErrTooBigValue
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String(), nil
}

type NumberType int

const (
	NumberRoman = NumberType(0)
	NumberArabic = NumberType(1)
)

func readNumber (s string) (int, NumberType, error) {
	if (isRoman(s[0])) {
		value, err := fromRoman(s)
		if (err != nil) {
			return 0, NumberArabic, err
		}
		return value, NumberRoman, nil
	} else {
		value, err := strconv.Atoi(s)

		if (err != nil) {
			return 0, NumberArabic, err
		}
		return value, NumberArabic, nil
	}
}

func writeNumber(n int, numbertype NumberType) bool {
	if (numbertype == NumberRoman) {
		str, err := toRoman(n)

		if (err != nil) {
			fmt.Println(err)
			return false
		}

		fmt.Println(str)
	} else {
		fmt.Println(n)
	}

	return true
}

type Operation int

var (
	ErrInvalidOperation = errors.New("Invalid operation")
)

const (
	Addition = Operation(0)
	Substruction = Operation(1)
	Multiplication = Operation(2)
	Division = Operation(3)
)

func readOp(s string) (Operation, error) {
	if (len(s)>1 || len(s)<1) {
		return 0, ErrInvalidOperation
	}
	switch (s[0]) {
	case '+':
		return Addition, nil
	case '-':
		return Substruction, nil
	case '*':
		return Multiplication, nil
	case '/':
		return Division, nil
	default:
		return 0, ErrInvalidOperation
	}
}

func main () {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')

		if (err != nil) {
			fmt.Println(err)
			break
		}

		// Remove newline if there is
		if text[len(text)-1] == '\n' {
			text = text[0:len(text)-1]
		}

		words := strings.Split(text, " ")

		if len(words) < 3 {
			fmt.Println("Too few arguments")
			break
		}

		if len(words) > 3 {
			fmt.Println("Too many arguments")
			break
		}

		for _, s := range(words) {
			if len(s) < 1 {
				fmt.Println("Empty string!")
				break
			}
		}

		val1, numbertype1, err := readNumber(words[0])

		if (err != nil) {
			fmt.Println(err)
			break
		}

		op, err := readOp(words[1])

		if (err != nil) {
			fmt.Println(err)
			break
		}

		val2, numbertype2, err := readNumber(words[2])

		if (err != nil) {
			fmt.Println(err)
			break
		}

		if (numbertype1 != numbertype2) {
			fmt.Println("Different types of numbers")
			break
		}

		var result int

		switch (op) {
		case Addition:
			result = val1 + val2
			break
		case Substruction:
			result = val1 - val2
			break
		case Multiplication:
			result = val1 * val2
			break
		case Division:
			result = val1 / val2
			break
		}

		if (!writeNumber(result, numbertype1)) {
			break
		}
	}
}
