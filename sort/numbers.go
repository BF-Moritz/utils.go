package sort

import (
	"sort"
)

func NumSortAsc[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](arr []T) []T {
	out := make([]T, 0, len(arr))
	copy(out, arr)

	sort.Slice(out, func(i, j int) bool {
		return out[i] > out[j]
	})
	return out
}

func NumSortDesc[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](arr []T) []T {
	out := make([]T, 0, len(arr))
	copy(out, arr)

	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}

func NumSortAscInPlace[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](arr []T) {
	sort.Slice(arr, func(i, j int) bool {
		return (arr)[i] < (arr)[j]
	})
}

func NumSortDescInPlace[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](arr []T) {
	sort.Slice(arr, func(i, j int) bool {
		return (arr)[i] > (arr)[j]
	})
}
