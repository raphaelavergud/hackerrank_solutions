package arrays

func removeDuplicates(array []int) int {

	// readPointer scans to find the next element
	// writePointer shows where the next non-duplicate int will be placed

	var writePointer int = 1

	for readPointer := 1; readPointer < len(array); readPointer++ {
		if array[readPointer] != array[writePointer-1] { // because the previous position is unique
			array[writePointer] = array[readPointer]
			writePointer++
		}
	}

	return writePointer
}
