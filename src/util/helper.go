package util

import "math"

func CalcIndexJob(length, nJob int) [][]int {
	step := math.Round(float64(length) / float64(nJob))

	result := [][]int{}
	for i := 0; i < length; i += int(step) {

		end := i + int(step) - 1
		if end > length {
			end = length
		}

		result = append(result, []int{i, end})
	}

	return result
}
