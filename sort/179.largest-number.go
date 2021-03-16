package sort

import (
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	rst := ""
	// In descending order
	s := MergeSort(nums)
	if s[0] == 0 {
		return "0"
	}
	for i := 0; i < len(s); i++ {
		rst += strconv.Itoa(s[i])
	}
	return rst
}
func cCompare(a, b int) bool {
	as, bs := strconv.Itoa(a), strconv.Itoa(b)
	return strings.Compare(as+bs, bs+as) == -1
}
func MergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return cmerge(left, right)
}

func cmerge(left []int, right []int) []int {
	newArr := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for {
		if cCompare(left[i], right[j]) {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}

		} else {
			newArr[index] = left[i]
			index++
			i++
			if i == len(left) {
				copy(newArr[index:], right[j:])
				break
			}
		}
	}
	return newArr
}
