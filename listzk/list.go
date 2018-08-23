package listzk

import (
	"errors"
)

//单链表
/*1.判断是否为空的单链表
2.单链表的长度
3.获取头节点
4.从头部添加元素
5.从尾部添加元素
6.在指定位置添加元素
7.删除指定元素
8.删除指定位置的元素
9.判断是否包含某个元素*/

//节点
type Node struct {
	Data interface{}
	Next *Node
}

//头节点
type List struct {
	HeadNode *Node
}

//判断是否为空
func (this *List) IsEmpty() bool {
	if this.HeadNode == nil {
		return true
	}
	return false
}

//获取链表的长度
func (this *List) Length() int {
	head := this.HeadNode
	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	return count
}

//获取头节点
func (this *List) GetHeadNode() *Node {
	return this.HeadNode
}

//从头节点插入
func (this *List) Add(x interface{}) {
	node := &Node{Data: x}
	node.Next = this.HeadNode
	this.HeadNode = node
}

//在尾部插入

func (this *List) Append(x interface{}) {
	node := &Node{Data: x}
	if this.IsEmpty() {
		this.HeadNode = node
	}
	head := this.HeadNode
	for head.Data != nil {
		head = head.Next
	}
	head.Next = node
}

//根据索引插入链表
func (this *List) Insert(index int, x interface{}) {
	if index < 0 {
		this.Add(x)
	} else if index > this.Length() {
		this.Append(x)
	} else {
		pre := this.HeadNode
		count := 0
		for count < (index - 1) {
			pre = pre.Next
			count++
		}
		node := &Node{Data: x}
		node.Next = pre.Next
		pre.Next = node
	}
}

//删除元素
func (this *List) Remove(x interface{}) {
	pre := this.HeadNode
	if pre.Data == x {
		this.HeadNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == x {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

//删除指定元素
func (this *List) RemoveAtIndex(index int) error {
	pre := this.HeadNode
	if index <= 0 {
		this.HeadNode = pre.Next
	} else if index > this.Length() {
		return errors.New("索引越界！")
	} else {
		count := 0
		if count < (index-1) && pre.Next != nil {
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next

	}
	return nil
}

//元素是否存在
func (this *List) Contain(x interface{}) bool {
	head := this.HeadNode
	for head.Data != nil {
		if head.Data == x {
			return true
		}
		head = head.Next
	}
	return false
}
