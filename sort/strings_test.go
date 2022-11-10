package sort_test

import (
	"testing"

	"github.com/BF-Moritz/utils.go/sort"
)

func TestStrSortAsc(t *testing.T) {

	unsortedArray := []string{"c", "test", "haha", "lol", "aha", "wtf", "hi"}
	sortedArray := []string{"aha", "c", "haha", "hi", "lol", "test", "wtf"}

	res := sort.StrSortAsc(unsortedArray)

	for i, v := range res {
		if v != sortedArray[i] {
			t.Fail()
		}
	}
}
func TestStrSortDesc(t *testing.T) {

	unsortedArray := []string{"c", "test", "haha", "lol", "aha", "wtf", "hi"}
	sortedArray := []string{"wtf", "test", "lol", "hi", "haha", "c", "aha"}

	res := sort.StrSortDesc(unsortedArray)

	for i, v := range res {
		if v != sortedArray[i] {
			t.Fail()
		}
	}
}
func TestStrSortAscInPlace(t *testing.T) {

	unsortedArray := []string{"c", "test", "haha", "lol", "aha", "wtf", "hi"}
	sortedArray := []string{"aha", "c", "haha", "hi", "lol", "test", "wtf"}

	sort.StrSortAscInPlace(unsortedArray)

	for i, v := range unsortedArray {
		if v != sortedArray[i] {
			t.Fail()
		}
	}
}
func TestStrSortDescInPlace(t *testing.T) {

	unsortedArray := []string{"c", "test", "haha", "lol", "aha", "wtf", "hi"}
	sortedArray := []string{"wtf", "test", "lol", "hi", "haha", "c", "aha"}

	sort.StrSortDescInPlace(unsortedArray)

	for i, v := range unsortedArray {
		if v != sortedArray[i] {
			t.Fail()
		}
	}
}
