package array_test

import (
	"testing"

	"github.com/BF-Moritz/utils.go/array"
)

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	f := func(v int) int {
		return v * 2
	}

	arrt := array.Map(arr, f)

	for i, a := range arrt {
		if a != arr[i]*2 {
			t.Fail()
		}
	}

	arr2 := []string{"das", "ist", "ein", "toller", "test"}

	f2 := func(v string) int {
		return len(v)
	}

	arrt2 := array.Map(arr2, f2)

	for i, v := range []int{3, 3, 3, 6, 4} {
		if v != arrt2[i] {
			t.Fail()
		}
	}
}
