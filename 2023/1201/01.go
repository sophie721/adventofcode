package _1201

import (
	"errors"
)

func Run01One(input []string) (int, error) {
	sum := 0
	for _, s := range input {
		first, err := findDigit(s, true)
		if err != nil {
			return 0, err
		}

		second, err := findDigit(s, false)
		if err != nil {
			return 0, err
		}

		sum += first*10 + second
	}
	return sum, nil
}

func findDigit(s string, lookFromHead bool) (int, error) {
	i, end, step := setUpIterator(len(s), lookFromHead)

	for i != end {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i]) - int('0'), nil
		}
		i += step
	}

	return 0, errors.New("Invalid input string")
}

func setUpIterator(strLen int, lookFromHead bool) (int, int, int) {
	var start, end, step int
	if lookFromHead {
		start = 0
		end = strLen
		step = 1
	} else {
		start = strLen - 1
		end = -1
		step = -1
	}
	return start, end, step
}
