package arrays

// merge two sorted slices into the first slice
// params are: nums1 & m: slice and number of elements in slice
// nums2 & n: slice and number of elements in slice
func merge(nums1 []int, m int, nums2 []int, n int) {
	var pointerForNums1 int = m - 1
	var pointerForNums2 int = n - 1

	// start from the end the writeIndex because there will be empty space, no need to replace existing items in the front of slice.
	for writeIndex := m + n - 1; pointerForNums2 >= 0; writeIndex-- {
		// check at end of nums1 if theres still an element there, and if there is and its bigger than the last one in nums2, place it in
		if pointerForNums1 >= 0 && nums1[pointerForNums1] > nums2[pointerForNums2] {
			nums1[writeIndex] = nums1[pointerForNums1]
			pointerForNums1 = pointerForNums1 - 1
		} else {
			// else, place element of nums2 in slice nums1
			nums1[writeIndex] = nums2[pointerForNums2]
			pointerForNums2 = pointerForNums2 - 1
		}
	}
}
