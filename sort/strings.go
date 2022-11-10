package sort

import (
	"sort"
	"strings"
)

func StrSortAsc(arr []string) []string {
	out := make([]string, 0, len(arr))
	copy(out, arr)

	sort.Slice(out, func(i, j int) bool {
		return strings.Compare((out)[i], (out)[j]) < 0
	})
	return out
}

func StrSortDesc(arr []string) []string {
	out := make([]string, 0, len(arr))
	copy(out, arr)

	sort.Slice(out, func(i, j int) bool {
		return strings.Compare((out)[i], (out)[j]) > 0
	})
	return out
}

func StrSortAscInPlace(arr []string) {
	sort.Slice(arr, func(i, j int) bool {
		return strings.Compare((arr)[i], (arr)[j]) < 0
	})
}

func StrSortDescInPlace(arr []string) {
	sort.Slice(arr, func(i, j int) bool {
		return strings.Compare((arr)[i], (arr)[j]) > 0
	})
}
