package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sumOfTails []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumOfTails = append(sumOfTails, 0)
		} else {
			sumOfTails = append(sumOfTails, Sum(numbers[1:]))
		}

	}

	return sumOfTails
}
