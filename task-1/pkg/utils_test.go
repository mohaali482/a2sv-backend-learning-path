package pkg_test

import (
	"testing"

	"github.com/mohaali482/a2sv-backend-learning-path/task-1/pkg"
)

func TestCalculateAverage(t *testing.T) {
	t.Run("Testing CalculateAverage with empty value", func(t *testing.T) {
		res := pkg.CalculateAverage(&map[string]int{})
		if res != 0 {
			t.Errorf("Expected zero result but found %v", res)
		}
	})
	t.Run("Testing CalculateAverage with some values", func(t *testing.T) {
		res := pkg.CalculateAverage(&map[string]int{
			"Math":    100,
			"English": 100,
			"Physics": 97,
		})
		if res != 99 {
			t.Errorf("Expected 99 result but found %v", res)
		}
	})
}
