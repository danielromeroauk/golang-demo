package sum

func Integers(numbers []int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
