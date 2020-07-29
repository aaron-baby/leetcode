| #        | Title                                                    | Solution                                    |
| -------- | -------------------------------------------------------- | ------------------------------------------- |
| 62.      | [Unique Paths][unique-paths]                             |   [Go](unique_paths.go)                     |
| 221.     | [Maximal Square][maximal-square]                         |   [Go](maximal_square.go)                   |
| 279.     | [Perfect Squares][perfect-squares]                       |   [Go](perfect_squares.go)                  |
| 416.     | [Partition Equal Subset Sum][partition-equal-subset-sum] |   [Go](partition_equal_subset_sum.go)       |

[unique-paths]: https://leetcode.com/problems/unique-paths/
[maximal-square]: https://leetcode.com/problems/maximal-square/
[perfect-squares]: https://leetcode.com/problems/perfect-squares/
[partition-equal-subset-sum]: https://leetcode.com/problems/partition-equal-subset-sum/

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
