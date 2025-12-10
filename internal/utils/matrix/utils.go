package matrix

import "math"

func GenerateBinarySequences(n int) Matrix {
	numSequences := int(math.Pow(2, float64(n)))
	result := NewEmptyMatrix(numSequences, n)
	for i := range numSequences {
		sequence := make([]int, n)

		temp := i

		// Fill the sequence array from right-to-left (least significant bit first)
		for j := n - 1; j >= 0; j-- {
			sequence[j] = temp & 1

			temp >>= 1
		}

		result.Rows[i].Values = sequence
	}

	return result
}
