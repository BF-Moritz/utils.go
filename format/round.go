package format

import (
	"fmt"
	"math"
	"strings"
)

func Round[T float32 | float64](num T, n int) string {
	str := fmt.Sprintf(fmt.Sprintf("%%.0%df", n), num)
	for strings.HasSuffix(str, "0") {
		str = strings.TrimSuffix(str, "0")
	}
	str = strings.TrimSuffix(str, ".")

	return str
}

func RoundToSignificant[T float32 | float64](num T, n int) string {

	mask := math.Pow10(n - 1)
	c := 0
	for float64(num) < mask {
		c--
		num *= 10
	}

	for float64(num) > (mask * 10) {
		c++
		num /= 10
	}

	num = T(math.Round(float64(num)))
	num *= T(math.Pow10(c))

	str := fmt.Sprintf(fmt.Sprintf("%%.0%df", n), num)
	for strings.HasSuffix(str, "0") {
		str = strings.TrimSuffix(str, "0")
	}
	str = strings.TrimSuffix(str, ".")

	return str
}
