package main

import "fmt"

func main() {
	array := []int{0,0,0}
	fmt.Println(threeSumClosest(array,1))
	// a := 0
	// ret := (a ^ a>>31) - a>>31
	// fmt.Println(ret,a)
}

func threeSumClosest(nums []int, target int) int {
	result := 0
    init := 0
    for ii,i := range nums {
        for jj,j:=range nums[ii+1:]{
            for _,l := range nums[ii+jj+2:]{
				sum := i+j+l
				suma := sum - target
				resa := result - target
				retsum := (suma ^ suma>>31) - suma>>31
				retres := (resa ^ resa>>31) - resa>>31
                if init == 0{
					result = sum
                    init++
                }
                if retsum < retres {
                    result = sum
                }
            }
        }
    }
    return result
}