package arrays

// Sum produces the sum of an array of integers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
