| #       | Title                                                    | Solution                                    |
| ------- | -------------------------------------------------------- | ------------------------------------------- |
| 2.      | [Add Two Numbers][add-two-numbers]                       |   [Go](add-two-numbers.go)                  |
| 19.     | [Remove Nth Node From End of List][remove-nth-node]      |   [Go](remove-nth-node.go)                  |
| Easy    | 21. [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)        |   [Go](merge_two_sorted_lists.go) |
| Hard    | 23. [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)            |   [Go](merge_k_sorted_lists.go)   |
| Medium  | 138. [Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/)     |   [Go](138.copy-list-with-random-pointer.go)   |
| Easy    | 141. [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)                 |   [Go](linked_list_cycle.go)      |
| Medium  | 142. [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)           |   [Go](linked-list-cycle-ii.go)   |
| Medium  | 143. [Reorder List](https://leetcode.com/problems/reorder-list/)                           |   [Go](reorder_list.go)           |
| Medium  | 147. [Insertion Sort List](https://leetcode.com/problems/insertion-sort-list/)             |   [Go](147.insertion-sort-list.go)           |
| Medium  | 148. [Sort List](https://leetcode.com/problems/sort-list/)                                 |   [Go](148.sort-list.go)                     |
| Easy    | 206. [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)             |   [Go](reverse_linked_list.go)    |
| Easy    | 234. [Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)       |   [Go](palindrome_linked_list.go) |
| Easy    | 876. [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) |   [Go](876.middle-of-the-linked-list.go)     |

[add-two-numbers]: https://leetcode.com/problems/add-two-numbers/
[remove-nth-node]: https://leetcode.com/problems/remove-nth-node-from-end-of-list/

## 21. Merge Two Sorted Lists
Similar Questions

[88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)

## 143. Reorder List
解法:
1. 先使用快慢指针找到中点
2. 反转后半段链表
3. 依次从前后链表中插值

## 148. Sort List
**解法：**
递归用快慢指针找到中点，分别排序左右链表，再合并两个有序链表