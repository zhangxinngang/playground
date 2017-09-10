package main

import (
	"strings"
)

func checkInclusion(s1 string, s2 string) bool {
	return strings.Contains(s2, s1) || strings.Contains(s2, reverseString(s1))
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
