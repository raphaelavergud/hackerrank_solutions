package arrays

func reverseArray(a []int32) []int32 {
	left := 0
	right := len(a) - 1

	for left < right {
		a[left], a[right] = a[right], a[left]

		left++
		right--
	}

	return a
}
