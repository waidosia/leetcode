# 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
# b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）： 
# 
#  
#  0 <= a, b, c, d < n 
#  a、b、c 和 d 互不相同 
#  nums[a] + nums[b] + nums[c] + nums[d] == target 
#  
# 
#  你可以按 任意顺序 返回答案 。 
# 
#  
# 
#  示例 1： 
# 
#  
# 输入：nums = [1,0,-1,0,-2,2], target = 0
# 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
#  
# 
#  示例 2： 
# 
#  
# 输入：nums = [2,2,2,2,2], target = 8
# 输出：[[2,2,2,2]]
#  
# 
#  
# 
#  提示： 
# 
#  
#  1 <= nums.length <= 200 
#  -10⁹ <= nums[i] <= 10⁹ 
#  -10⁹ <= target <= 10⁹ 
#  
#  Related Topics 数组 双指针 排序 👍 1037 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
from typing import List


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        l = len(nums)
        result = []
        nums.sort()
        for j in range(l):
            if j > 0 and nums[j] == nums[j - 1]:
                continue
            for i in range(j + 1, l):
                left, right = i + 1, l - 1
                if i > j + 1 and nums[i - 1] == nums[i]:
                    continue
                while left < right:
                    _ = nums[i] + nums[left] + nums[right] + nums[j]
                    if _ == target:
                        result.append([nums[j], nums[i], nums[left], nums[right]])
                        while left != right and nums[left] == nums[left + 1]:
                            left += 1
                        while left != right and nums[right] == nums[right - 1]:
                            right -= 1

                        left += 1
                        right -= 1
                    if _ > target:
                        right -= 1
                    if _ < target:
                        left += 1

        return result


# leetcode submit region end(Prohibit modification and deletion)