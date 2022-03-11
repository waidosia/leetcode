//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的
//功能。(若队列中没有元素，deleteHead 操作返回 -1 ) 
//
// 
//
// 示例 1： 
//
// 输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
// 
//
// 示例 2： 
//
// 输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
// 
//
// 提示： 
//
// 
// 1 <= values <= 10000 
// 最多会对 appendTail、deleteHead 进行 10000 次调用 
// 
// Related Topics 栈 设计 队列 👍 452 👎 0

package main

import "container/list"

//leetcode submit region begin(Prohibit modification and deletion)
type CQueue struct {
	stack1,stack2 *list.List
}


func Constructor() CQueue {
	return CQueue{
		stack1:list.New(),
		stack2:list.New(),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stack1.PushBack(value)
}


func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len()>0{
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0{
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)
