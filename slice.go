package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := a[5:7]
	c := append(a, 1)
	fmt.Println("len,cap:c", len(c), cap(c))

	fmt.Println("len,cap", len(b), cap(b))
	b = append(b, 2)
	fmt.Println("len,cap", len(b), cap(b))
	b = append(b, 3)
	fmt.Println("len,cap", len(b), cap(b))
	b = append(b, 4)
	fmt.Println("len,cap", len(b), cap(b))
	fmt.Println(b, a)
	a = append(a, 1)
	fmt.Println("len,cap ,a", len(a), cap(a))
	a = append(a, []int{1, 2, 3, 4, 5, 5, 6, 9}...)
	fmt.Println("len,cap ,a", len(a), cap(a))
	// r := []int{2, 0}
	// merge(r, 1, []int{1}, 1)
	// fmt.Println(r)
}
