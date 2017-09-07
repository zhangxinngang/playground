package main

import (
	"fmt"
)

type NumMatrix struct {
	Elements [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{Elements: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row, col := 0, 0
	sum := 0
	for row = row1; row <= row2; row++ {
		for col = col1; col <= col2; col++ {
			sum += this.Elements[row][col]
		}
	}
	return sum
}

func main() {
	mx := [][]int{
		[]int{3, 0, 1, 4, 2},
		[]int{5, 6, 3, 2, 1},
		[]int{1, 2, 0, 1, 5},
		[]int{4, 1, 0, 1, 7},
		[]int{1, 0, 3, 0, 5}}
	matrix := Constructor(mx)
	fmt.Println(matrix.SumRegion(1, 2, 2, 4))
}
