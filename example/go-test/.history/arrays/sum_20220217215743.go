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
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumTails(numbersToSum ...[]int) ([]int) {
	
	var sums []int

	for _, 
}