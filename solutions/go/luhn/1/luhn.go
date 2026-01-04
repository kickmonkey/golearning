package luhn

import "unicode"

func Valid(id string) bool {
	sum := 0
	countDigits := 0
	double := false // 从右往左：第一位不翻倍，第二位翻倍...

	for i := len(id) - 1; i >= 0; i-- {
		r := rune(id[i])

		if r == ' ' {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}

		d := int(r - '0')
		countDigits++

		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}

	return countDigits > 1 && sum%10 == 0
}