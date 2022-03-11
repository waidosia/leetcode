//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
// 示例 1：
//
// 输入：s = "We are happy."
//输出："We%20are%20happy."
//
//
//
// 限制：
//
// 0 <= s 的长度 <= 10000
// Related Topics 字符串 👍 206 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	// 统计空格个数
	count := 0
	for _, i := range s {
		if i == ' ' {
			count++
		}
	}

	b := []byte(s)
	length := len(b)
	tmp := make([]bytvvve, count*2)
	b = append(b, tmp...)
	i, j := length-1, len(b)-1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}

	}
	return string(b)
}

//leetcode submit region end(Prohibit modification and deletion)
