func myAtoi(s string) int {
	i := 0
	n := len(s)

	// Spaces
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	// sign
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		sign = 1
		i++
	}

	// digits
	result := 0
	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		//overflow
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++

	}

	return result * sign
}
