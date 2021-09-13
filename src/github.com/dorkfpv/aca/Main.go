package main

import (
	"fmt"
)

const (
	rowCount = 2
)

func main() {
	matrixA := [rowCount]int{}
	matrixB := []int{}
	matrixB = append(matrixB, 1, 2, 3)
	fmt.Printf("%v %T\n", matrixB, matrixB)
	fmt.Printf("%v %T\n", matrixA, matrixA)
	fmt.Printf("%v \n", cap(matrixB))
}
