package main

func validPalindrome(s string) bool {
	var i, j =0, len(s)-1
	for i<j{
		if s[i] == s[j]{
			i++
			j--
		} else {
			// remove char at position i or j
			return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j])
		}
	}

	return true
}

func isPalindrome(s string) bool {
	var i, j =0, len(s)-1
	for i<j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}