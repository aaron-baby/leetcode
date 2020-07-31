> The height of a node is the length of the longest downward path to a leaf from that node. 
> 
> The height of the root is the height of the tree.

| Diffct   | Title                                      | Solution                           |
| -------- | ------------------------------------------ | ---------------------------------- |
| Medium   | 94. [Binary Tree Inorder Traversal][binary-tree-inorder-traversal]              | [Go](inorder_travesal.go)             |
| Medium   | 103. [Binary Tree Zigzag Level Order Traversal][zigzag-level-order-traversal]   | [Go](binary_tree_zigzag_travelsal.go) |
| Easy     | 226. [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)    | [Go](invert_binary_tree.go)           |
| Medium   | 230. [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)    | [Go](kth_smallest_element_in_a_bst.go)           |
| Easy     | 543. [Diameter of Binary Tree][diameter-of-binary-tree]                         | [Go](diameter_of_binary_tree.go)      |

[binary-tree-inorder-traversal]: https://leetcode.com/problems/binary-tree-inorder-traversal/
[zigzag-level-order-traversal]: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
[diameter-of-binary-tree]: https://leetcode.com/problems/diameter-of-binary-tree/

## 94. Binary Tree Inorder Traversal
in-order traversal retrieves the keys in ascending sorted order.

## 230. Kth Smallest Element in a BST
这个解法使用二叉树中序遍历和堆栈，好处是不用构建整个树，找到最小的第k个元素就返回了