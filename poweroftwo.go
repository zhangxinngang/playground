package main

import "fmt"

func isPowerOfTwo(n int) bool {
	//判断n是不是2的次方
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n < 0 {
		return false
	}
	mod := n % 2
	if mod != 0 {
		return false
	}
	for mod == 0 && n > 2 {
		n = n / 2
		mod = n % 2
		fmt.Println("n,mod:", n, mod)
		if mod != 0 {
			return false
		}
	}
	return true
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}

func main() {
	//fmt.Println(isPowerOfTwo(96))
	fmt.Println(isPowerOfTwo2(16))
}
