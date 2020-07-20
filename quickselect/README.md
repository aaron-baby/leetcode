> In computer science, quickselect is a selection algorithm to find the kth smallest element in an unordered list.

pseudo code
```
QuickSelect(A, k)
  let r be chosen uniformly at random in the range 1 to length(A)
  let pivot = A[r]
  let A1, A2 be new arrays
  # split into a pile A1 of small elements and A2 of big elements
  for i = 1 to n
    if A[i] < pivot then
      append A[i] to A1
    else if A[i] > pivot then
      append A[i] to A2
    else
      # do nothing
  end for
  if k <= length(A1):
    # it's in the pile of small elements
    return QuickSelect(A1, k)
  else if k > length(A) - length(A2)
    # it's in the pile of big elements
    return QuickSelect(A2, k - (length(A) - length(A2))
  else
    # it's equal to the pivot
    return pivot
```

| #        | Title                                                    | Solution                                    |
| -------- | -------------------------------------------------------- | ------------------------------------------- |
| 215      | [Kth Largest Element in an Array](kth-largest-element)   |   [Go](kth_largest_element_in_an_array.go)  |


[kth-largest-element-in-an-array]: https://leetcode.com/problems/kth-largest-element-in-an-array/