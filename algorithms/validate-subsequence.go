package algorithms

func ValidateSubsequence(sequence, subSequence []int) bool {
	i, j := 0, 0
	for i < len(sequence) && j < len(subSequence) {
		if sequence[i] == subSequence[j] {
			j++
		}
		i++
	}
	return j == len(subSequence)
}
