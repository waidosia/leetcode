//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
//
//
//
// 示例 1：
//
//
//输入：timePoints = ["23:59","00:00"]
//输出：1
//
//
// 示例 2：
//
//
//输入：timePoints = ["00:00","23:59","00:00"]
//输出：0
//
//
//
//
// 提示：
//
//
// 2 <= timePoints <= 2 * 10⁴
// timePoints[i] 格式为 "HH:MM"
//
// Related Topics 数组 数学 字符串 排序 👍 134 👎 0

package main

import (
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func getMinutes(t string) int {
	// 获取当前时间的秒数
	return (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func findMinDifference(timePoints []string) int {
	if len(timePoints) > (60 * 24) {
		return 0
	}
	sort.Strings(timePoints)
	ans := math.MaxInt32
	t0Minutes := getMinutes(timePoints[0])
	preMinutes := t0Minutes
	for _, t := range timePoints[1:] {
		minutes := getMinutes(t)
		ans = min(minutes-preMinutes, ans)
		preMinutes = minutes
	}
	ans = min(ans,t0Minutes + 60*24 - preMinutes)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
