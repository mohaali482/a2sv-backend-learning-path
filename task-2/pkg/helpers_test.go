package pkg_test

import (
	"testing"

	"github.com/mohaali/a2sv-backend-learning-path/task-2/pkg"
)

func TestPalindromeFormatString(t *testing.T) {
	t.Run("Testing with no characters", func(t *testing.T) {
		str := pkg.PalindromeFormatString("")

		if str != "" {
			t.Errorf("Expected empty string, but found %v \n", str)
		}
	})

	t.Run("Testing with special characters", func(t *testing.T) {
		str := pkg.PalindromeFormatString("!@#$%^&*().?><:';][}{}]\\+_~`\"")

		if str != "" {
			t.Errorf("Expected empty string, but found %v \n", str)
		}
	})

	t.Run("Testing with no special characters", func(t *testing.T) {
		str := pkg.PalindromeFormatString("hello")

		if str != "hello" {
			t.Errorf("Expected hello, but found %v \n", str)
		}
	})

	t.Run("Testing with capitalized characters", func(t *testing.T) {
		str := pkg.PalindromeFormatString("HELLO")

		if str != "hello" {
			t.Errorf("Expected hello, but found %v \n", str)
		}
	})

	t.Run("Testing with capitalized characters and special characters", func(t *testing.T) {
		str := pkg.PalindromeFormatString("HELLO!@#$%^&*().?><:';][}{}]\\+_~`\"")

		if str != "hello" {
			t.Errorf("Expected hello, but found %v \n", str)
		}
	})

}
