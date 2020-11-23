> Backtracking is a general algorithm for finding all (or some) solutions to some computational problems, 
> 
> notably constraint satisfaction problems, that incrementally builds candidates to the solutions, and abandons a candidate ("backtracks") as soon as it determines that the candidate cannot possibly be completed to a valid solution.

| Diffct | Title                        | Solution                           |
| ------ | ---------------------------- | ---------------------------------- |
| Medium | 17. [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | [Go](letter_combinations.go)  |
| Medium | 22. [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)                                   | [Go](generate_parentheses.go) |
| Hard   | 37. [Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)                                                 | [Go](37.sudoku-solver.go)     |
| Medium | 39. [Combination Sum](https://leetcode.com/problems/combination-sum/)                                             | [Go](combination_sum.go)      |
| Medium | 40. [Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)                                       | [Go](40.combination-sum-ii.go)      |
| Medium | 47. [Permutations II](https://leetcode.com/problems/permutations-ii/)                                             | [Go](permutations-ii.go)            |
| Hard   | 51. [N-Queens](https://leetcode.com/problems/n-queens/)                                                           | [Go](n_queens.go)                   |
| Hard   | 52. [N-Queens II](https://leetcode.com/problems/n-queens-ii/)                                                     | [Go](n-queens-ii.go)                |
| Medium | 79. [Word Search](https://leetcode.com/problems/word-search/)                                                     | [Go](79.word-search.go)                        |
| Medium | 89. [Gray Code](https://leetcode.com/problems/gray-code/)                                                         | [Go](89.gray-code.go)                          |
| Hard   | 996. [Number of Squareful Arrays](https://leetcode.com/problems/number-of-squareful-arrays/)                      | [Go](996.number-of-squareful-arrays.go)        |

## 52. N-Queens II
Basically, we have to ensure 4 things:
1. No two queens share a column.
2. No two queens share a row.
3. No two queens share a top-right to left-bottom diagonal.
4. No two queens share a top-left to bottom-right diagonal.

**解题思路**

先在第0行上放置皇后，依次检查从0到n-1列有没有被占用，如果空置则把皇后放在该列并置占用
```
for col := 0; col < n; col++ {
    if notUnderAttack(row, col, n, rows, hills, dales) {
        rows[col] = 1
    }
}
```
山丘对角线`/`方向上的row+col=常数，e.g. row 0+col 0 =0
(1, 0), (0, 1)的行列号和等于1

山谷对角线`\`的行号-列号等于常数[1-n, n-1]，加2n是为了避免负值的下标