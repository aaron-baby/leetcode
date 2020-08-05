package linked_list

func mergeKLists(lists []*ListNode) *ListNode {
	amount := len(lists)
	interval := 1
	for interval < amount {
		for i := 0; i < amount-interval; i += interval * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	if amount > 0 {
		return lists[0]
	}
	return nil
}
