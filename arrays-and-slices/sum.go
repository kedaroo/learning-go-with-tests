package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	// this is slower 👇🏻, yields 38 ns/op in benchmarks
	// var results []int

	// this is faster 👇🏻, yield 20 ns/ops in benchmarks
	/*
	 * this is faster because slice capacity is set at the time of its initialization
	 */
	results := make([]int, 0, len(slices))

	for _, slice := range slices {
		var sum int

		for _, val := range slice {
			sum += val
		}

		results = append(results, sum)
	}

	return results
}

func SumAllTails(slices ...[]int) []int {
	results := make([]int, len(slices))

	for i, slice := range slices {
		if len(slice) > 0 {
			results[i] = Sum(slice[1:])
		} else {
			results[i] = 0
		}
	}

	return results
}
