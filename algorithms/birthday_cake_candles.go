package algorithms

func birthdayCakeCandles(candles []int32) int32 {

	var tallestHeight int32
	tallestHeight = candles[0]

	var countOfTallest int32
	countOfTallest = 1

	var index int
	index = 1

	for index < len(candles) {
		var currentCandleHeight int32
		currentCandleHeight = candles[index]

		if currentCandleHeight > tallestHeight {
			tallestHeight = currentCandleHeight
			countOfTallest = 1

		} else if currentCandleHeight == tallestHeight {
			countOfTallest = countOfTallest + 1
		}

		index = index + 1
	}
	return countOfTallest
}
