package format

import "fmt"

func FSBase10[T int | int16 | int32 | int64 | float32 | float64 | uint | uint16 | uint32 | uint64](num T) string {
	c := 0
	rem := float64(num)
	for rem > 1000 {
		rem /= 1000
		c++
	}

	arr := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	return fmt.Sprintf("%s%s", RoundToSignificant(rem, 3), arr[c])
}

func FSBase2[T int | int16 | int32 | int64 | float32 | float64 | uint | uint16 | uint32 | uint64](num T) string {
	c := 0
	rem := float64(num)
	for rem > 1024 {
		rem /= 1024
		c++
	}

	arr := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}

	return fmt.Sprintf("%02f%s", rem, arr[c])
}
