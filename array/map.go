package array

func Map[T, U int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64 | string](arr []T, function func(T) U) []U {
	out := make([]U, len(arr))
	for i := 0; i < len(arr); i++ {
		out[i] = function(arr[i])
	}
	return out
}
