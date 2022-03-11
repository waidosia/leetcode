//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//
// 示例 2：
//
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//
//
// 提示：
//
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
// Related Topics 链表 👍 1167 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	// 找到left 的上一个节点
	leftNodePrev := dummy
	for i := 0; i < left-1; i++ {
		leftNodePrev = leftNodePrev.Next
	}
	// 找到right节点
	rightNode := leftNodePrev
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	// 切断字链表
	leftNode := leftNodePrev.Next
	rightNodeNext := rightNode.Next
	leftNodePrev.Next = nil
	rightNode.Next = nil
	// 反转子链表
	leftNode = reverseList(leftNode)
	// 重新连接
	leftNodePrev.Next = leftNode
	for leftNode.Next != nil{
		leftNode = leftNode.Next
	}
	leftNode.Next = rightNodeNext
	return dummy.Next
}


func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	if head == nil {
		return nil
	}
	next := head.Next
	for next != nil {
		cur.Next = prev
		prev = cur
		cur = next
		next = next.Next
	}
	cur.Next = prev
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
