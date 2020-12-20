package main

import (
	"fmt"
	"github.com/intrntsrfr/datastructures"
)

func main() {

	l := []int{ 1, 3, 4, 4, 6, 17, 79, 81, 90}
	a := 79

	fmt.Println(datastructures.BinarySearch(l, a))
}