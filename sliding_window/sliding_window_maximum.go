package sliding_window

import (
	"container/list"
)

func maxSlidingWindow(nums []int, k int) []int {
	var dequeue = list.New()
	var res []int
	for i := 0; i < k; i++ {
		removeSmaller(dequeue, nums, i)
		dequeue.PushBack(i)
	}
	// The element at the front of the queue is the largest element of previous window
	e := dequeue.Front()
	res = append(res, nums[e.Value.(int)])
	for i := k; i < len(nums); i++ {
		if dequeue.Front() != nil && dequeue.Front().Value.(int) <= i-k {
			dequeue.Remove(dequeue.Front())
		}
		removeSmaller(dequeue, nums, i)
		dequeue.PushBack(i)
		fe := dequeue.Front()
		res = append(res, nums[fe.Value.(int)])
	}

	return res
}

func removeSmaller(dequeue *list.List, nums []int, i int) {
	for dequeue.Back() != nil && nums[i] >= nums[dequeue.Back().Value.(int)] {
		dequeue.Remove(dequeue.Back())
	}
}
