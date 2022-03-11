//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 4388 👎 0

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left,right := i + 1,len(nums) - 1
		for left < right{
			n2,n3 := nums[left],nums[right]
			if n1 + n2 + n3 == 0{
				res = append(res,[]int{n1,n2,n3})
				for left < right && nums[left] == n2 {
					left ++
				}
				for left < right && nums[right] == n3 {
					right --
				}
			} else if n1 +n2 +n3 <0 {
				left ++
			} else {
				right --
			}
		}

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
