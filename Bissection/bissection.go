package bissection

import (
	"fmt"
	"math"
)

func Calculate(number float64) float64 {

	calculated := math.Sqrt(math.Pow(number-1, 2)+1)/(number-1) - 10/number

	return calculated
}

func Range(i float64, f float64) (float64, float64) {
	a := Calculate(i)

	b := Calculate(f)

	m := (i + f) / 2

	fmt.Println(m)

	c := Calculate(m)

	if a > 0 && b > 0 {
		return a, b
	}

	if a > 0 && c < 0 || a < 0 && c > 0 {
		return i, m
	} else if b > 0 && c < 0 || b < 0 && c > 0 {
		return f, m
	} else {
		return i, f
	}

}
