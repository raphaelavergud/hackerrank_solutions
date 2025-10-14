package arrays

func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0
	var index int = 0
	for index < len(ar) {
		var currentNumber int32 = ar[index]
		sum = sum + currentNumber
		index = index + 1
	}
	return sum
}
