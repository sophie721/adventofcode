package _1201

import (
	"errors"
	"fmt"
	"slices"
	"sync"
)

func Run01Two(input []string) (int, error) {
	sum := 0
	wg := new(sync.WaitGroup)
	var lock sync.Mutex

	addNumber := func(s string, sum *int, errs []error) {
		first, err := findDigitAndString(s, true)
		if err != nil {
			lock.Lock()
			errs = append(errs, err)
			lock.Unlock()
		}

		second, err := findDigitAndString(s, false)
		if err != nil {
			lock.Lock()
			errs = append(errs, err)
			lock.Unlock()
		}

		lock.Lock()
		fmt.Printf("%s: %d + %d\n", s, *sum, (first*10 + second))
		*sum = *sum + first*10 + second
		lock.Unlock()

		wg.Done()
	}

	var errs []error
	wg.Add(len(input))
	for _, s := range input {
		go addNumber(s, &sum, errs)
	}
	wg.Wait()

	if len(errs) > 0 {
		fmt.Printf("Error occured: %s\n", errs[0].Error())
		return 0, errs[0]
	}

	return sum, nil
}

func findDigitAndString(s string, lookFromHead bool) (int, error) {
	var buffer []rune
	start, end, step := setUpIterator(len(s), lookFromHead)
	for i := start; i != end; i += step {
		if s[i] >= 'a' && s[i] <= 'z' {
			buffer = append(buffer, rune(s[i]))
			result, err := letter2Int(buffer, !lookFromHead)
			if err == nil {
				return result, nil
			}
			continue
		}

		buffer = buffer[:0]

		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i]) - int('0'), nil
		}
		break
	}

	return 0, errors.New("Invalid input")
}

func letter2Int(b []rune, reverse bool) (int, error) {
	digitLetters := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	bb := make([]rune, len(b))
	copy(bb, b)
	if reverse {
		slices.Reverse(bb)
	}
	str := string(bb)
	len := len(str)
	for i := 0; i < len; i++ {
		// letter min length = 3
		for j := i + 3; j <= len; j++ {
			for idx, letter := range digitLetters {
				if letter == str[i:j] {
					return idx + 1, nil
				}
			}
		}
	}

	return 0, errors.New("No letter matched.")
}
