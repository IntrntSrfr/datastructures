package main

import (
	"fmt"
	"math/big"
)

func main() {
	/*
		l := []int{ 1, 3, 4, 4, 6, 17, 79, 81, 90}
		a := 79

		fmt.Println(datastructures.BinarySearch(l, a))
	*/

	fmt.Println(BigFactorial(big.NewInt(400)))
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func BigFactorial(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1)
	}
	res := new(big.Int).Set(n)
	return res.Mul(res, BigFactorial(n.Sub(n, big.NewInt(1))))
	//return n.Mul(n, BigFactorial(n.Sub(n, big.NewInt(1))))
}
