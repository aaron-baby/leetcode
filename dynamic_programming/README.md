| #        | Title                                                    | Solution                                    |
| -------- | -------------------------------------------------------- | ------------------------------------------- |
| Medium   | 5. [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)       |   [Go](5.longest-palindromic-substring.go)       |
| Hard     | 44. [Wildcard Matching](https://leetcode.com/problems/wildcard-matching/)                              |   [Go](44.wildcard-matching.go)                  |
| Medium   | 62.[Unique Paths](https://leetcode.com/problems/unique-paths/)                                         |   [Go](unique_paths.go)                     |
| Medium   | 63. [Unique Paths II](https://leetcode.com/problems/unique-paths-ii/)                                  |   [Go](63.unique-paths-ii.go)                |
| Medium   | 64. [Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/)                                |   [Go](64.minimum-path-sum.go)               |
| Medium   | 91. [Decode Ways](https://leetcode.com/problems/decode-ways/)                                                |   [Go](91.decode-ways.go)                         |
| Easy     | 122. [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) |   [Go](122.best-time-to-buy-and-sell-stock-ii.go) |
| Medium   | 152. [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)                     |   [Go](152.maximum-product-subarray.go)       |
| Medium   | 198. [House Robber](https://leetcode.com/problems/house-robber/)                                             |   [Go](198.house-robber.go)                   |
| Medium   | 213. [House Robber II](https://leetcode.com/problems/house-robber-ii/)                                       |   [Go](213.house-robber-ii.go)                |
| Medium   | 221. [Maximal Square](https://leetcode.com/problems/maximal-square/)                                         |   [Go](maximal_square.go)                     |
| Medium   | 279. [Perfect Squares](https://leetcode.com/problems/perfect-squares/)                                       |   [Go](perfect_squares.go)                    |
| Medium   | 309. [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)                   |   [Go](309.best-time-to-buy-and-sell-stock-with-cooldown.go)  |
| Hard     | 403. [Frog Jump](https://leetcode.com/problems/frog-jump/)   |   [Go](frog_jump.go)                          |
| Medium   | 416. [Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/)                 |   [Go](partition_equal_subset_sum.go)         |
| Medium   | 714. [Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)     |   [Go](714.best-time-to-buy-and-sell-stock-with-transaction-fee.go) |


## 279. Perfect Squares
计算每一个小于n的平方根组合

状态转移方程

    temp := x * x
    dp[i] = Min(dp[i], 1+dp[i-temp])

## 416. Partition Equal Subset Sum
维护一个二维数组，行是从0到half sum，列是从空集依次增加一个元素到整个数组。
p(i, j)表示是否能从集合{ x1, ..., xj }里找到一个子集的和等于sum
如果前一个子集{ x1, ..., x<sub>j-1</sub> }为真当前子集也是真值

> there is a subset of { x1, ..., xj−1 } that sums to i − xj, since xj + that subset's sum = i.

或者{ x1, ..., x<sub>j−1</sub> }子集的和等于i-x<sub>j</sub>，加上x<sub>j</sub>刚好和等于i
> [1, 5, 11, 5]

这个例子里的half sum是11，比如说求p(6,1)，要么p(6,0)为真或者p(6-1,)

|     | {} <sub>j=0</sub>  | {1} <sub>j=1</sub> | {1,5} <sub>j=2</sub> | {1,5,11} <sub>j=3</sub> | {1,5,11,5} <sub>j=4</sub> |
| --- | --- | --- | ----- | -------- | ---------- |
| 0   |     |     |       |          |            |
| 1   |     |     |       |          |            |
| 2   |     |     |       |          |            |
| 3   |     |     |       |          |            |
| 4   |     |     |       |          |            |
| 5   |     |     |       |          |            |
| 6   |     |     |       |          |            |
| 7   |     |     |       |          |            |
| 8   |     |     |       |          |            |
| 9   |     |     |       |          |            |
| 10  |     |     |       |          |            |
| 11  |     |     |       |          |            |

## 91. Decode Ways
**解题思路**

如果最后一个数字是1到9，说明该数位可以被转成字符，状态转移为前i-1个数字的解法。

最后两个数位的值在10到26之间，说明两个数字的组合可行，累加dp[i-2]的解法