package algorithms

import "sort"

func TwoNumberSum(array []int, target int) []int {
	keysValues := make(map[int]int)
	for i, item := range array {
		checkNumber := target - item
		if j, exists := keysValues[checkNumber]; exists {
			return []int{j, i}
		}
		keysValues[item] = i
	}
	return []int{}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
