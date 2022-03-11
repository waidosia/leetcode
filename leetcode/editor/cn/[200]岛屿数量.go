//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 1571 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) (result int) {
	row, col := len(grid), len(grid[0])
	// 辅助数组
	visited := make([][]int, row)
	for i := range visited {
		visited[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' && visited[i][j] == 0 {
				dfs(grid, i, j, visited)
				result++
			}
		}
	}
	return
}

func dfs(grid [][]byte, i int, j int, visited [][]int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' || visited[i][j] == 1 {
		return
	}
	visited[i][j] = 1
	dfs(grid, i-1, j, visited)
	dfs(grid, i+1, j, visited)
	dfs(grid, i, j-1, visited)
	dfs(grid, i, j+1, visited)
}

//leetcode submit region end(Prohibit modification and deletion)
