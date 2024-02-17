package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	var results []int

	for _, slice := range slices {
		var sum int

		for _, val := range slice {
			sum += val
		}

		results = append(results, sum)
	}
	
	return results
}
