package statistics

import (
	"math"

	"gonum.org/v1/gonum/stat/combin"
)

func Slovin(n int, e float64) int {
	fn := float64(n)
	return int(math.Ceil(fn / (1 + fn*e*e)))
}

func PossibleHandCount(pn int) int {
	return combin.Binomial(52, (pn<<1)+5)
}
