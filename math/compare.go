package math

func Min[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](a T, b T) T {
	if a > b {
		return b
	}
	return a
}

func Max[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](a T, b T) T {
	if a < b {
		return b
	}
	return a
}

func MinArray[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](arr []T) T {
	if len(arr) <= 0 {
		return T(0)
	}

	var min T = arr[0]

	for _, val := range arr {
		if val < min {
			min = val
		}
	}

	return min
}

func MaxArray[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](arr []T) T {
	if len(arr) <= 0 {
		return T(0)
	}

	var max T = arr[0]

	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}
