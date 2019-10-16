package main

import (
	. "github.com/isdamir/gotype" //引入定义的数据结构
)

//带头结点的逆序
//就地逆序,time O(n),space O(1)
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

//递归逆序
func RecursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func RecursiveReverse(node *LNode) {
	firstNode := node.Next
	//递归调用
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

//插入法
func InsertReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *LNode
	var next *LNode
	cur = node.Next.Next
	//设置链表的第一个节点为尾节点
	node.Next.Next = nil
	//把遍历到的节点插入到头节点后面
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

func main() {

}
