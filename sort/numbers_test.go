package sort_test

import (
	"testing"

	"github.com/BF-Moritz/utils.go/sort"
)

func TestNumSortAsc(t *testing.T) {

	unsortedArray := []int{1, 7, 3, -10, 213, 3, 69, 420, 11}
	sortedArray := []int{-10, 1, 3, 3, 7, 11, 69, 213, 420}

	res := sort.NumSortAsc(unsortedArray)

	for i, v := range res {
		if v != sortedArray[i] {
			t.Fail()
		}
	}
}
func TestNumSortDesc(t *testing.T) {

	unsortedArray := []int{1, 7, 3, -10, 213, 3, 69, 420, 11}
	sortedArray := []int{420, 213, 69, 11, 7, 3, 3, 1, -10}

	res := sort.NumSortDesc(unsortedArray)

	for i, v := range res {
		if v != sortedArray[i] {
			t.Fail()
		}
	}
}
func TestNumSortAscInPlace(t *testing.T) {

	unsortedArray := []int{1, 7, 3, -10, 213, 3, 69, 420, 11}
	sortedArray := []int{-10, 1, 3, 3, 7, 11, 69, 213, 420}

	sort.NumSortAscInPlace(unsortedArray)

	for i, v := range unsortedArray {
		if v != sortedArray[i] {
			t.Fail()
		}
	}
}
func TestNumSortDescInPlace(t *testing.T) {

	unsortedArray := []int{1, 7, 3, -10, 213, 3, 69, 420, 11}
	sortedArray := []int{420, 213, 69, 11, 7, 3, 3, 1, -10}

	sort.NumSortDescInPlace(unsortedArray)

	for i, v := range unsortedArray {
		if v != sortedArray[i] {
			t.Fail()
		}
	}
}
