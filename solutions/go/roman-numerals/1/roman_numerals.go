package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("invalid number")
	}

	// 从大到小排列，贪心匹配
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	n := input
	out := make([]byte, 0, 16)

	for i, v := range values {
		for n >= v {
			n -= v
			out = append(out, symbols[i]...)
		}
	}

	return string(out), nil
}
