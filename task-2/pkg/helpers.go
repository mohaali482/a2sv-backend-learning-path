package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringInput() string {
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')

	return strings.TrimSpace(inp)
}

func GetIntInput() int {
	input := StringInput()
	num, err := strconv.Atoi(input)

	for err != nil {
		fmt.Println(err)
		fmt.Println("Please try again.")
		fmt.Print("Insert value: ")
		input = StringInput()
		num, err = strconv.Atoi(input)
	}

	return num
}

func validRuneForPalindrome(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}

	if r >= 'a' && r <= 'z' {
		return true
	}

	return false
}

func validRuneForCounter(r rune) bool {
	if r == ' ' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}

	if r >= 'a' && r <= 'z' {
		return true
	}

	return false
}

func PalindromeFormatString(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	strBuilder := strings.Builder{}

	for _, r := range str {
		if validRuneForPalindrome(r) {
			strBuilder.WriteRune(r)
		}
	}

	return strBuilder.String()
}

func CounterFormatString(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	strBuilder := strings.Builder{}

	for _, r := range str {
		if validRuneForCounter(r) {
			strBuilder.WriteRune(r)
		}
	}

	return strBuilder.String()
}
