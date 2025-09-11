package arrays

func hourglassSum(arr [][]int32) int32 {
	var maxHourglassSum int32 = -100

	row := 0
	for row <= 3 {
		column := 0
		for column <= 3 {
			sumTop := arr[row][column] + arr[row][column+1] + arr[row][column+2]

			sumMiddle := arr[row+1][column+1]

			sumBottom := arr[row+2][column] + arr[row+2][column+1] + arr[row+2][column+2]

			currentHourglassSum := sumTop + sumMiddle + sumBottom

			if currentHourglassSum > maxHourglassSum {

				maxHourglassSum = currentHourglassSum
			}

			column = column + 1
		}
		row = row + 1
	}
	return maxHourglassSum

}
