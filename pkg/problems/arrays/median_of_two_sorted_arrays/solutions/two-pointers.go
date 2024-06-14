package arrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0

	// While both arrays still have values, keep looping through
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i += 1
		} else {
			nums3 = append(nums3, nums2[j])
			j += 1
		}
	}

	// Append any remaining bits from the longer array that are missing
	if i < len(nums1) {
		nums3 = append(nums3, nums1[i:]...)
	}
	if j < len(nums2) {
		nums3 = append(nums3, nums2[j:]...)
	}

	// Calc median of the new array
	mid := len(nums3) / 2
	if len(nums3)%2 == 0 {
		// Even numbers. Mid index = 2 in [10,20,30,40], average of 2 & 1 index values equals median
		return (float64(nums3[mid]) + float64(nums3[mid-1])) / 2.0
	} else {
		// Odd numbers. Literally just the middle index
		return float64(nums3[mid])
	}
}
