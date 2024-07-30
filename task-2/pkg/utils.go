package pkg

import (
	"strings"
)

func CountWords(str string) int {
	str = strings.TrimSpace(str)
	counter := 0
	wordLength := 0

	for i := range len(str) {
		if str[i] == ' ' {
			if wordLength > 0 {
				counter++
				wordLength = 0
			}
		} else {
			wordLength++
		}
	}

	if wordLength > 0 {
		counter++
	}

	return counter
}
