package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := [5]int{1, 3, 5, 6, 7}
	nums2 := [3]int{2, 10, 12}

	expectedMedian := 5.5
	actualMedian := FindMedianSortedArrays(nums1[:], nums2[:])

	if actualMedian != expectedMedian {
		t.Fatalf(`FindMedianSortedArrays(nums1[:], nums2[:]) = %f, want %f, error`, actualMedian, expectedMedian)
	}
}
