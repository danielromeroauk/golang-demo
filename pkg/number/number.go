package number

const Billion int = 1000000000

func GetSliceInt(length int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = i
	}
	return numbers
}
