package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := a[5:7]
	b = append(b, 2)
	fmt.Println(b, a)
}
