//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
// 
//
// 示例 2： 
//
// 
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 10 
// -100 <= matrix[i][j] <= 100 
// 
// Related Topics 数组 矩阵 模拟 👍 1001 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	res := make([]int,0)
	for  i := 0;i<len(matrix[0])/2+1;i++{
		// 上
		for j := i ; j<len(matrix[0])-i-1 && len(matrix) -1 -i >= i; j++ {
			res = append(res,matrix[i][j])
		}
		// 右
		for j := i;j < len(matrix)-i && i < len(matrix[0])-1-i;j++{
			res = append(res,matrix[j][len(matrix)-i-1])
		}
		// 下
		for j := len(matrix[0])-i-2;j>i && len(matrix) -1 -i>i; j-- {
			res = append(res, matrix[len(matrix)-1-i][j])
		}
		// 左
		for j := len(matrix) - 1 - i; j > i && i < len(matrix[0])-1-i; j-- {
			res = append(res, matrix[j][i])
		}

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
