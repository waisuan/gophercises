package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, num := range numbersToSum {
		sums = append(sums, Sum(num))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, num := range numbersToSum {
		if len(num) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := num[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
