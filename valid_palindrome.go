package main

import "fmt"

func validPalindrome(s string) bool {
	return recursive_check(s, 0)
}

func recursive_check(s string, c int) bool {
	l := len(s)
	fmt.Println(s)
	if l == 1 {
		return true
	} else if l == 2 {
		return s[0] == s[1]
	}
	if c > 1 {
		return false
	}
	// compare first and last
	if s[0] == s[l-1] {
		return recursive_check(s[1:l-1], c)
	} else {
		c += 1
		return recursive_check(s[1:l], c) || recursive_check(s[0:l-1], c)
	}
}
func main() {
	//fmt.Println(validPalindrome("aba"))
	//fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abcad"))
}
