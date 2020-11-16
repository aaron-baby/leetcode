package hash_table

func twoSum(nums []int, target int) []int {
	// hash that remember element index
	indices := make(map[int]int)
	for i, v := range nums {
		difference := target - v
		j, ok := indices[difference]
		if ok {
			return []int{j, i}
		} else {
			indices[v] = i
		}
	}

	return []int{}
}
