package math

import (
	"fmt"
	"math"
	"strconv"
)

// MustParsePrecFloat64 按小数位截取float
func MustParsePrecFloat64(value float64, prec int) float64 {
	format := fmt.Sprintf("%%.%df", prec)
	v, err := strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	if err != nil {
		return 0
	}
	return v
}

func ChangeFloat64toInt(value float64) int {
	relizeValue := math.Floor(value)
	return int(relizeValue)
}
