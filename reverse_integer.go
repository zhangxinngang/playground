package main

import (
	"fmt"
	"strconv"
)

func Reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res < -2147483648 || res > 2147483647 {
		return 0
	}
	return res
}

func Sqrt(x int) int {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 1
	case 4:
		return 2
	}
	sqrt := 0
	l := 1
	h := x
	m := (l + h) / 2
	for {
		if m*m > x {
			h = m
		}
		if m*m < x {
			l = m
		}
		if m*m == x {
			return m
		}
		m = (h + l) / 2
		sqrt = m
		if h-l <= 1 {
			fmt.Println("h,l,m", h, l, m)
			break
		}
	}
	return sqrt
}

func compareVersion(version1 string, version2 string) int {
	v1, _ := strconv.ParseFloat(version1, 64)
	v2, _ := strconv.ParseFloat(version2, 64)
	fmt.Println(v1, v2)
	if v1 > v2 {
		return 1
	} else if v1 < v2 {
		return -1
	} else {
		return 0
	}
}

func firstMissingPositive(nums []int) int {
	intmap := map[int]int{}
	for _, v := range nums {
		intmap[v] = 1
	}
	i := 1
	for i < len(nums)+1 {
		v := intmap[i]
		if v == 0 {
			return i
		}
		i++
	}
	return i
}

var TelephoneMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	result := TelephoneMap[string(digits[0])]
	for i, letter := range digits {
		if i == 0 {
			continue
		}
		if nil == TelephoneMap[string(letter)] {
			break
		}
		result = CartesianProduct(result, TelephoneMap[string(letter)])
	}
	return result
}

func CartesianProduct(a []string, b []string) []string {
	result := []string{}
	for _, aa := range a {
		for _, bb := range b {
			result = append(result, aa+bb)
		}
	}
	return result
}

func intersection(nums1 []int, nums2 []int) []int {
	map1 := map[int]int{}
	map2 := map[int]int{}
	for _, v := range nums1 {
		map1[v] = 1
	}
	for _, v := range nums2 {
		map2[v] = 1
	}
	mm := map1
	nn := map2
	if len(map1) > len(map2) {
		mm = map2
		nn = map1
	}
	result := []int{}
	if len(mm) == 0 {
		return result
	}
	for k, _ := range mm {
		if nn[k] == 1 {
			result = append(result, k)
		}
	}
	return result
}

func Dist(p1 []int, p2 []int) int {
	return (p2[1]-p1[1])*(p2[1]-p1[1]) + (p2[0]-p1[0])*(p2[0]-p1[0])
}

func CheckEqual(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return Dist(p1, p2) > 0 && Dist(p1, p2) == Dist(p2, p3) &&
		Dist(p2, p3) == Dist(p3, p4) && Dist(p3, p4) == Dist(p4, p1) &&
		Dist(p1, p3) == Dist(p2, p4)
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return CheckEqual(p1, p2, p3, p4) || CheckEqual(p1, p3, p2, p4) || CheckEqual(p1, p2, p4, p3)
}

func main() {
	//fmt.Println(Reverse())
	//fmt.Println(Sqrt(7))
	//fmt.Println(compareVersion("1", "1"))
	//fmt.Println(firstMissingPositive([]int{1, 2, 3, 4}))
	//fmt.Println(letterCombinations("2a"))
	//fmt.Println(intersection([]int{4, 7, 9, 7, 6, 7}, []int{5, 0, 0, 6, 1, 6, 2, 2, 4}))
	fmt.Println(validSquare([]int{0, 0}, []int{5, 0}, []int{5, 4}, []int{0, 4}))
}
