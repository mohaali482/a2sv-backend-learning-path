package pkg_test

import (
	"testing"

	"github.com/mohaali/a2sv-backend-learning-path/task-2/pkg"
)

func TestCountWords(t *testing.T) {
	t.Run("Testing with no word", func(t *testing.T) {
		count := pkg.CountWords("")

		if len(count) != 0 {
			t.Errorf("Expected length count to be 0, but found %v \n", len(count))
		}
	})

	t.Run("Testing with one word", func(t *testing.T) {
		count := pkg.CountWords("One")

		if count["one"] != 1 {
			t.Errorf("Expected count of one to be 1, but found %v \n", count["one"])
		}
	})

	t.Run("Testing with one word and extra spaces on the left", func(t *testing.T) {
		count := pkg.CountWords("       One")

		if count["one"] != 1 {
			t.Errorf("Expected count of one to be 1, but found %v \n", count["one"])
		}
	})

	t.Run("Testing with one word and extra spaces on the right", func(t *testing.T) {
		count := pkg.CountWords("One       ")

		if count["one"] != 1 {
			t.Errorf("Expected count of one to be 1, but found %v \n", count["one"])
		}
	})

	t.Run("Testing with one word and extra spaces on the both sides", func(t *testing.T) {
		count := pkg.CountWords("      One       ")

		if count["one"] != 1 {
			t.Errorf("Expected count of one to be 1, but found %v \n", count["one"])
		}
	})

	t.Run("Testing with two word and extra spaces on the middle", func(t *testing.T) {
		count := pkg.CountWords("One       One")

		if count["one"] != 2 {
			t.Errorf("Expected count of one to be 2, but found %v \n", count["one"])
		}
	})

	t.Run("Testing with two word and extra spaces on the middle and on the right", func(t *testing.T) {
		count := pkg.CountWords("One       One            ")

		if count["one"] != 2 {
			t.Errorf("Expected count of one to be 2, but found %v \n", count["one"])
		}
	})

	t.Run("Testing with two word and extra spaces on the middle and on the left", func(t *testing.T) {
		count := pkg.CountWords("            One       One")

		if count["one"] != 2 {
			t.Errorf("Expected count of one to be 2, but found %v \n", count["one"])
		}
	})

	t.Run("Testing with two word and extra spaces on the middle and on the both sides", func(t *testing.T) {
		count := pkg.CountWords("            One       One            ")

		if count["one"] != 2 {
			t.Errorf("Expected count of one to be 2, but found %v \n", count["one"])
		}
	})

	t.Run("Testing with two word and extra spaces on the middle and on the both sides", func(t *testing.T) {
		count := pkg.CountWords("            One       Two            ")

		if count["one"] != 1 || count["two"] != 1 {
			t.Errorf("Expected count of one to be 1 and count of two to be 1, but found %v, %v \n", count["one"], count["two"])
		}
	})
}

func TestPalindromeCheck(t *testing.T) {
	t.Run("Testing with no characters", func(t *testing.T) {
		palindrome := pkg.PalindromeCheck("")

		if !palindrome {
			t.Errorf("Expected palindrome to be true, but found %v \n", palindrome)
		}
	})

	t.Run("Testing with odd length palindrome word", func(t *testing.T) {
		palindrome := pkg.PalindromeCheck("12321")

		if !palindrome {
			t.Errorf("Expected palindrome to be true, but found %v \n", palindrome)
		}
	})

	t.Run("Testing with even length palindrome word", func(t *testing.T) {
		palindrome := pkg.PalindromeCheck("1221")

		if !palindrome {
			t.Errorf("Expected palindrome to be true, but found %v \n", palindrome)
		}
	})

	t.Run("Testing with even length palindrome word with special characters", func(t *testing.T) {
		palindrome := pkg.PalindromeCheck("12!@#$21")

		if !palindrome {
			t.Errorf("Expected palindrome to be true, but found %v \n", palindrome)
		}
	})

	t.Run("Testing with even length non-palindromic word", func(t *testing.T) {
		palindrome := pkg.PalindromeCheck("125321")

		if palindrome {
			t.Errorf("Expected palindrome to be false, but found %v \n", palindrome)
		}
	})

	t.Run("Testing with odd length non-palindromic word", func(t *testing.T) {
		palindrome := pkg.PalindromeCheck("1254321")

		if palindrome {
			t.Errorf("Expected palindrome to be false, but found %v \n", palindrome)
		}
	})
}
