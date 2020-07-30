package array

func findDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, cnt := low+(high-low)>>1, 0
		for _, item := range nums {
			if item <= mid {
				cnt++
			}
		}
		// not duplicate numbers less than mid should just equal to mid
		if cnt > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
