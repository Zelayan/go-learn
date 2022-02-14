package arrays

func Sum(numbers []int) (sum int) {
	sum = 0

	for _, v := range numbers {
		sum += v
	}
	return 
}

func SumAll(numbersToSum ...[]int) ([]int) {
	var sums []int
	for i, numbers := range numbersToSum {
		sums = Sum(numbers)
	}
	return sums
}