package algorithms_test

import (
	"testing"

	"github.com/i3onilha/go-learning/algorithms"
)

func TestValidateSubsequence(t *testing.T) {
	scenarios := []struct {
		array       []int
		subsequence []int
		expected    bool
	}{
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 6, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{22, 25, 6}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{25}, true},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1}, true},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1}, true},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}, true},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1}, true},
		{[]int{1, 1, 1, 1, 1}, []int{1}, true},
		{[]int{1, 1, 1, 1, 1}, []int{2}, false},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10, 12}, false},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{4, 5, 1, 22, 25, 6, -1, 8, 10}, false},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{-1, 6, 10}, false},
	}
	for i, scenario := range scenarios {
		result := algorithms.ValidateSubsequence(scenario.array, scenario.subsequence)
		if result != scenario.expected {
			t.Errorf("Expected %v, but got %v on scenario %d", scenario.expected, result, i)
		}
	}
}
