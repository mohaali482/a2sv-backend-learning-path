package pkg

import "strings"

func CountWords(str string) map[string]int {
	str = CounterFormatString(str)
	counter := make(map[string]int)
	word := strings.Builder{}

	for _, r := range str {
		if r == ' ' {
			counter[word.String()]++
			word.Reset()
		} else {
			word.WriteRune(r)
		}
	}
	if word.Len() > 0 {
		counter[word.String()]++
	}

	return counter
}

func PalindromeCheck(str string) bool {
	str = PalindromeFormatString(str)

	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
