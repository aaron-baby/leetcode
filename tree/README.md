> The height of a node is the length of the longest downward path to a leaf from that node. 
> 
> The height of the root is the height of the tree.

| Diffct   | Title                                      | Solution                           |
| -------- | ------------------------------------------ | ---------------------------------- |
| Medium   | 94. [Binary Tree Inorder Traversal][binary-tree-inorder-traversal]              | [Go](inorder_travesal.go)             |
| Medium   | 98. [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)                           | [Go](98.validate-binary-search-tree.go)             |
| Easy     | 101. [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)            | [Go](101.symmetric-tree.go)           |
| Medium   | 102. [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)              | [Java](/java/src/Solution.java)             |
| Medium   | 103. [Binary Tree Zigzag Level Order Traversal][zigzag-level-order-traversal]   | [Go](binary_tree_zigzag_travelsal.go) |
| Easy     | 226. [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)    | [Go](invert_binary_tree.go)           |
| Medium   | 230. [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)                      | [Go](kth_smallest_element_in_a_bst.go)  |
| Medium   | 236. [Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)  | [Go](lca_of_a_binary_tree.go)           |
| Easy     | 501. [Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)                  | [Go](find_mode.go)<br>[TypeScript](src/tree/find_mode.ts)               |
| Easy     | 543. [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/)                                  | [Go](diameter_of_binary_tree.go)        |

[binary-tree-inorder-traversal]: https://leetcode.com/problems/binary-tree-inorder-traversal/
[zigzag-level-order-traversal]: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

## 94. Binary Tree Inorder Traversal
in-order traversal retrieves the keys in ascending sorted order.

## 230. Kth Smallest Element in a BST
这个解法使用二叉树中序遍历和堆栈，好处是不用构建整个树，找到最小的第k个元素就返回了

## 236. Lowest Common Ancestor of a Binary Tree
> In graph theory and computer science, the lowest common ancestor (LCA) of two nodes v and w in a tree or directed acyclic graph (DAG) T 
> is the lowest (i.e. deepest) node that has both v and w as descendants

## 501. Find Mode in Binary Search Tree
> In computer science, a binary search tree (BST), also called an ordered or sorted binary tree, 
>
> is a rooted binary tree whose internal nodes each store a key **greater than all the keys in the node's left subtree** and less than those in its right subtree.