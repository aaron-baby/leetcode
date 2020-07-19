package union_find

// 128. Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	appearance := make(map[int]bool)
	for _, v := range nums {
		appearance[v] = true
	}
	var ans int
	for i := 0; i < len(nums); i++ {
		// current number be the first element of sequence
		_, ok := appearance[nums[i]-1]
		if !ok {
			j := nums[i] + 1
			for {
				_, ok1 := appearance[j]
				if !ok1 {
					if j-nums[i] > ans {
						ans = j - nums[i]
					}
					break
				}
				j++
			}
		}
	}

	return ans
}
