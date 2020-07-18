package backtracking

// 17. Letter Combinations of a Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var output []string
	if len(digits) > 0 {
		backtrack("", digits, &output, m)
	}
	return output
}
func backtrack(combination string, next_digits string, slicePtr *[]string, m map[string]string) {
	if len(next_digits) == 0 {
		output := *slicePtr
		*slicePtr = append(output, combination)
		return
	}
	digit := next_digits[:1]
	letters := m[digit]
	for i := 0; i < len(letters); i++ {
		l := letters[i : i+1]
		backtrack(combination+l, next_digits[1:], slicePtr, m)
	}
}
