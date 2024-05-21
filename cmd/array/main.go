package main

import "fmt"

func main() {
	nums1 := [5]int{1, 3, 5, 6, 7}
	nums2 := [3]int{2, 10, 12}

	m := FindMedianSortedArrays(nums1[:], nums2[:])
	fmt.Printf("Median: %f", m)
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m, n := len(nums1), len(nums2)

	// A new array to store the sorted list (-2 to account for the len() starting at 1)
	var nums3 []int = make([]int, m+n-2)

	// Pointers for the current position in each list
	pos, i, j := 0, 0, 0

	// While nums1/2 pointer is less than nums1/2 length, keep looping
	for j < n && i < m {
		if pos >= len(nums3) {
			panic("We didnt calc the length of the nums3 array very well did we...")
		}

		head1, head2 := nums1[i], nums2[j]

		if head1 > head2 {
			nums3[pos] = head2
			j += 1
		} else {
			nums3[pos] = head1
			i += 1
		}

		pos += 1
	}

	// Append any remaining bits from the longer array that are missing
	if j < n {
		nums3 = append(nums3, nums2[j:]...)
	}

	if i < m {
		nums3 = append(nums3, nums1[i:]...)
	}

	var median float64
	if len(nums3)%2 == 0 {
		// even
		mIdx := len(nums3) / 2
		median = (float64(nums3[mIdx]) + float64(nums3[mIdx-1])) / float64(2)
	} else {
		// odd
		mIdx := len(nums3) / 2
		median = float64(nums3[mIdx])
	}

	return median
}
