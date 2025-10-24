package algorithms

func diagonalDifference(arr [][]int32) int32 {
	// top-left to bottom-right
	var primaryDiagonalSum int32
	primaryDiagonalSum = 0

	// top-right to bottom-left
	var secondaryDiagonalSum int32
	secondaryDiagonalSum = 0

	// number of rows in the given square matrix
	var matrixSize int
	matrixSize = len(arr)

	var rowIndex int
	rowIndex = 0

	for rowIndex < matrixSize {
		// for this diagonal row index and column index are the same
		// 0 0, 1 1, 2 2
		var primaryElement int32
		primaryElement = arr[rowIndex][rowIndex]

		primaryDiagonalSum = primaryDiagonalSum + primaryElement

		// for this diagonal: column index is matrixSize - 1 - rowIndex cuz it goes bottom to left top
		// like this 0 2, 1 1, 2 0 in small 3x3matrix
		var secondaryElement int32
		var secondaryColumnIndex int
		secondaryColumnIndex = matrixSize - 1 - rowIndex
		secondaryElement = arr[rowIndex][secondaryColumnIndex]

		secondaryDiagonalSum = secondaryDiagonalSum + secondaryElement
		// end secondary

		rowIndex = rowIndex + 1
	}

	var difference int32
	difference = primaryDiagonalSum - secondaryDiagonalSum
	if difference < 0 {
		return difference * -1
	}
	return difference
}
