package algorithms_test

import (
	"testing"

	"github.com/i3onilha/go-learning/algorithms"
)

func TestTwoNumberSum(t *testing.T) {
	scenarios := []struct {
		array  []int
		target int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5}, 6, []int{1, 3}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{2, 3}},
		{[]int{1, 2, 3, 4, 5}, 8, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{0, 2}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{}},
		{[]int{1, 2, 3, 4, 5}, -1, []int{}},
	}
	for i, scenario := range scenarios {
		result := algorithms.TwoNumberSum(scenario.array, scenario.target)
		if !algorithms.Equal(result, scenario.result) {
			t.Errorf("Expected %v, got %v on index %d", scenario.result, result, i)
		}
	}
}
