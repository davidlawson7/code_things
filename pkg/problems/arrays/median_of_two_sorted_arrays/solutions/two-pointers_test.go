package arrays

import "testing"

func TestFindMedianSortedArrays1(t *testing.T) {
	nums1 := [5]int{1, 3, 5, 6, 7}
	nums2 := [3]int{2, 10, 12}

	expectedMedian := 5.5
	actualMedian := FindMedianSortedArrays(nums1[:], nums2[:])

	if actualMedian != expectedMedian {
		t.Fatalf(`FindMedianSortedArrays(nums1[:], nums2[:]) = %f, want %f, error`, actualMedian, expectedMedian)
	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	nums1 := [2]int{1, 3}
	nums2 := [1]int{2}

	expectedMedian := 2.0
	actualMedian := FindMedianSortedArrays(nums1[:], nums2[:])

	if actualMedian != expectedMedian {
		t.Fatalf(`FindMedianSortedArrays(nums1[:], nums2[:]) = %f, want %f, error`, actualMedian, expectedMedian)
	}
}

func TestFindMedianSortedArrays3(t *testing.T) {
	nums1 := [2]int{1, 2}
	nums2 := [2]int{3, 4}

	expectedMedian := 2.5
	actualMedian := FindMedianSortedArrays(nums1[:], nums2[:])

	if actualMedian != expectedMedian {
		t.Fatalf(`FindMedianSortedArrays(nums1[:], nums2[:]) = %f, want %f, error`, actualMedian, expectedMedian)
	}
}

func TestFindMedianSortedArraysNums1Empty(t *testing.T) {
	nums1 := [0]int{}
	nums2 := [2]int{3, 4}

	expectedMedian := 3.5
	actualMedian := FindMedianSortedArrays(nums1[:], nums2[:])

	if actualMedian != expectedMedian {
		t.Fatalf(`FindMedianSortedArrays(nums1[:], nums2[:]) = %f, want %f, error`, actualMedian, expectedMedian)
	}
}

func TestFindMedianSortedArraysNums2Empty(t *testing.T) {
	nums1 := [5]int{1, 2, 3, 4, 5}
	nums2 := []int{}

	expectedMedian := 3.0
	actualMedian := FindMedianSortedArrays(nums1[:], nums2[:])

	if actualMedian != expectedMedian {
		t.Fatalf(`FindMedianSortedArrays(nums1[:], nums2[:]) = %f, want %f, error`, actualMedian, expectedMedian)
	}
}
