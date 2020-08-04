> A sliding window is an abstract concept commonly used in array/string problems. 

> A window is a range of elements in the array/string which usually defined by the start and end indices, i.e. [i,j) (left-closed, right-open). 

> A sliding window is a window "slides" its two boundaries to the certain direction.

| Diffct   | Title                                                    | Solution                                    |
| -------- | -------------------------------------------------------- | ------------------------------------------- |
| Medium   | 3. [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) |   [Go](longest_substring.go)  |
| Hard     | 76. [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)   |   [Go](minimum_window_substring.go)    |
| Hard     | 239. [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)      |   [Go](sliding_window_maximum.go)      |

## 239. Sliding Window Maximum
解题思路，使用deque双端队列
> In computer science, a double-ended queue (abbreviated to deque, pronounced _deck_) 
> is an abstract data type that generalizes a queue, for which elements can be added to or removed from either the front (head) or back (tail).