package Newton

import (
	"math"
)

func fX(number float64) float64 {

	calculated := math.Pow(number, 2) - 2

	return calculated
}

func fPrimeX(number float64) float64 {

	calculated := 2 * number

	return calculated

}

func X(xI float64) float64 {
	fx := fX(xI)

	fpx := fPrimeX(xI)

	xF := xI - (fx / fpx)

	return xF
}
