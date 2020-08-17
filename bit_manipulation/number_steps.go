package bit_manipulation

// 1404. Number of Steps to Reduce a Number in Binary Representation to One
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
func numSteps(s string) int {
	if len(s) == 1 {
		return 0
	}
	var count int
	for i := len(s) - 1; i > 0; {
		// even number, divide by 2, right shift
		if s[i] == '0' {
			count++
			i--
			continue
		}
		// add 1 operation turn trailing `1`s to `0` e.g. 10(11)->11(00)
		count++
		// trailing `0`s means it's an even number, right shift
		for s[i] == '1' && i > 0 {
			count++
			i--
		}
		// last bit
		if i == 0 {
			count++
		}
		// set higher bit to 1
		s = s[:i] + "1" + s[i+1:]
	}

	return count
}
