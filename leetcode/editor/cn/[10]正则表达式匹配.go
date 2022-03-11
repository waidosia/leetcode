//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
//
//
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
//
// 示例 1：
//
//
//输入：s = "aa", p = "a"
//输出：false
//解释："a" 无法匹配 "aa" 整个字符串。
//
//
// 示例 2:
//
//
//输入：s = "aa", p = "a*"
//输出：true
//解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//
//
// 示例 3：
//
//
//输入：s = "ab", p = ".*"
//输出：true
//解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// 1 <= p.length <= 30
// s 只包含从 a-z 的小写字母。
// p 只包含从 a-z 的小写字母，以及字符 . 和 *。
// 保证每次出现字符 * 时，前面都匹配到有效的字符
//
// Related Topics 递归 字符串 动态规划 👍 2779 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ { //s字符串可以为空串，所以从0开始
		for j := 1; j <= n; j++ { //p字符串不能为空串，所以从1开始
			if p[j-1] == '*' {
				if i != 0 && (s[i-1] == p[j-2] || p[j-2] == '.') { //p前一个字符与s的字符匹配
					dp[i][j] = dp[i][j-2] || dp[i-1][j] //匹配0次或匹配多次，匹配多次可以当作将s[i-1][j-1]舍弃掉
				} else { //不匹配，那么*前面字符被匹配0次
					dp[i][j] = dp[i][j-2]
				}
			} else {
				if i != 0 && (s[i-1] == p[j-1] || p[j-1] == '.') { //当前字符能匹配，则等于前一个字符匹配情况
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
