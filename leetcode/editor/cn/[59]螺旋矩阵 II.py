# 给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。 
# 
#  
# 
#  示例 1： 
# 
#  
# 输入：n = 3
# 输出：[[1,2,3],[8,9,4],[7,6,5]]
#  
# 
#  示例 2： 
# 
#  
# 输入：n = 1
# 输出：[[1]]
#  
# 
#  
# 
#  提示： 
# 
#  
#  1 <= n <= 20 
#  
#  Related Topics 数组 矩阵 模拟 👍 541 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
from typing import List


class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        matrix = [[0] * n for _ in range(n)]
        number = 1
        up, right, down, left = 0, n - 1, n - 1, 0
        while left < right and up < down:

            for x in range(left, right):
                matrix[up][x] = number
                number += 1

            for y in range(up, down):
                matrix[y][right] = number
                number += 1

            for x in range(right, left, -1):
                matrix[down][x] = number
                number += 1

            for y in range(down, up, -1):
                matrix[y][left] = number
                number += 1
            up += 1
            right -= 1
            down -= 1
            left += 1

        if n % 2:
            matrix[n // 2][n // 2] = number
        return matrix

# leetcode submit region end(Prohibit modification and deletion)
