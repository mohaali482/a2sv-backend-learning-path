package pkg_test

import (
	"testing"

	"github.com/mohaali/a2sv-backend-learning-path/task-2/pkg"
)

func TestCountWords(t *testing.T) {
	t.Run("Testing with no word", func(t *testing.T) {
		count := pkg.CountWords("")

		if count != 0 {
			t.Errorf("Expected count to be 0, but found %v \n", count)
		}
	})

	t.Run("Testing with one word", func(t *testing.T) {
		count := pkg.CountWords("One")

		if count != 1 {
			t.Errorf("Expected count to be 1, but found %v \n", count)
		}
	})

	t.Run("Testing with one word and extra spaces on the left", func(t *testing.T) {
		count := pkg.CountWords("       One")

		if count != 1 {
			t.Errorf("Expected count to be 1, but found %v \n", count)
		}
	})

	t.Run("Testing with one word and extra spaces on the right", func(t *testing.T) {
		count := pkg.CountWords("One       ")

		if count != 1 {
			t.Errorf("Expected count to be 1, but found %v \n", count)
		}
	})

	t.Run("Testing with one word and extra spaces on the both sides", func(t *testing.T) {
		count := pkg.CountWords("      One       ")

		if count != 1 {
			t.Errorf("Expected count to be 1, but found %v \n", count)
		}
	})

	t.Run("Testing with two word and extra spaces on the middle", func(t *testing.T) {
		count := pkg.CountWords("One       Two")

		if count != 2 {
			t.Errorf("Expected count to be 2, but found %v \n", count)
		}
	})

	t.Run("Testing with two word and extra spaces on the middle and on the right", func(t *testing.T) {
		count := pkg.CountWords("One       Two            ")

		if count != 2 {
			t.Errorf("Expected count to be 2, but found %v \n", count)
		}
	})

	t.Run("Testing with two word and extra spaces on the middle and on the left", func(t *testing.T) {
		count := pkg.CountWords("            One       Two")

		if count != 2 {
			t.Errorf("Expected count to be 2, but found %v \n", count)
		}
	})

	t.Run("Testing with two word and extra spaces on the middle and on the both sides", func(t *testing.T) {
		count := pkg.CountWords("            One       Two            ")

		if count != 2 {
			t.Errorf("Expected count to be 2, but found %v \n", count)
		}
	})
}
