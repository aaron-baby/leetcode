package sort

func pancakeSort(arr []int) []int {
	var res []int
	sp(arr, len(arr), &res)
	return res
}

func sp(arr []int, n int, resP *[]int) {
	if n == 1 {
		return
	}
	// 1. find max pancake index
	maxCake, maxIndex := 0, 0
	for i := 0; i < n; i++ {
		if arr[i] > maxCake {
			maxCake = arr[i]
			maxIndex = i
		}
	}
	// 2. -1 move max pancake to top
	reverse(arr, 0, maxIndex)
	*resP = append(*resP, maxIndex+1)
	//    -2 move max pancake to bottom
	reverse(arr, 0, n-1)
	*resP = append(*resP, n)

	// 3. recursive call previous n-1 pancake
	sp(arr, n-1, resP)
}

func reverse(cakes []int, i, j int) {
	for i < j {
		cakes[i], cakes[j] = cakes[j], cakes[i]
		i++
		j--
	}
}
