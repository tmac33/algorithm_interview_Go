package main

import (
	. "github.com/isdamir/gotype" //引入定义的数据结构
)

//带头结点的逆序
func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	//定义前驱节点
	var pre *LNode

	//定义当前节点
	var cur *LNode

	//后继节点
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

func main() {

}
