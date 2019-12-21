package main

import (
	"math/big"
)

func squareRoot(n int64, precision int64) int {

	limit := new(big.Int).Exp(big.NewInt(10), big.NewInt(precision+1), nil)

	a := big.NewInt(5 * n)
	b := big.NewInt(5)
	five := big.NewInt(5)
	ten := big.NewInt(10)
	hundred := big.NewInt(100)

	for b.Cmp(limit) < 0 {
		if a.Cmp(b) < 0 {
			a.Mul(a, hundred)
			tmp := new(big.Int).Div(b, ten)
			tmp.Mul(tmp, hundred)
			b.Add(tmp, five)
		} else {
			a.Sub(a, b)
			b.Add(b, ten)
		}
	}
	b.Div(b, hundred)

	ansDec := b.String()

	inc := 0
	for _, a := range ansDec {
		inc += int(a - '0')
	}

	return inc
}

func main() {
	println(squareRoot(2, 1000))
}
