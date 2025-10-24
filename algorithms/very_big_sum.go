package algorithms

// just a basic sum solution but this time with int64
// i googled and it can store numbers from:
// -9,223,372,036,854,775,808 to +9,223,372,036,854,775,807.

func aVeryBigSum(ar []int64) int64 {
	// int64 because result may be too big for normal int
	var totalSum int64
	totalSum = 0

	// index variable to keep track of position
	var index int
	index = 0

	for index < len(ar) {
		var currentNumber int64
		currentNumber = ar[index]

		totalSum = totalSum + currentNumber

		index = index + 1
	}

	return totalSum
}
