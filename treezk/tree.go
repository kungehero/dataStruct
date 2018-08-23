package treezk

import "fmt"

type Object interface{}

//队列二叉树
/*
0.追加元素
1.广度遍历
2.先序遍历
3.中序遍历
4.后序遍历*/
type Node struct {
	Data  Object
	LNode *Node
	RNode *Node
}

type Tree struct {
	RootNode *Node
}

//广度优先，按层级遍历添加
//追加元素 (广度优先，即按层级遍历后添加)
func (this *Tree) Add(object Object) {
	node := &Node{Data: object}
	if this.RootNode == nil {
		this.RootNode = node
		return
	}
	queue := []*Node{this.RootNode}
	for len(queue) != 0 {
		cur_node := queue[0]
		queue = queue[1:]

		if cur_node.LNode == nil {
			cur_node.LNode = node
			return
		} else {
			queue = append(queue, cur_node.LNode)
		}
		if cur_node.RNode == nil {
			cur_node.RNode = node
			return
		} else {
			queue = append(queue, cur_node.RNode)
		}
	}
}

//广度遍历
func (this *Tree) BreadthTravel() {

	if this.RootNode == nil {
		return
	}
	queue := []*Node{}
	queue = append(queue, this.RootNode)

	for len(queue) != 0 {
		//fmt.Printf("len(queue):%d\n", len(queue))
		cur_node := queue[0]
		queue = queue[1:]

		fmt.Printf("%v  ", cur_node.Data)

		if cur_node.LNode != nil {
			queue = append(queue, cur_node.LNode)
		}
		if cur_node.RNode != nil {
			queue = append(queue, cur_node.RNode)
		}
	}

}

//先序
func (this *Tree) PreOrder(node *Node) {
	if node == nil {
		return
	}
	println(&node.Data)
	this.PreOrder(node.LNode)
	this.PreOrder(node.RNode)
}

//中序
func (this *Tree) InOrder(node *Node) {
	if node == nil {
		return
	}
	this.PreOrder(node.LNode)
	println(&node.Data)
	this.PreOrder(node.RNode)
}

//后序
func (this *Tree) PostOrder(node *Node) {
	if node == nil {
		return
	}
	this.PreOrder(node.LNode)
	this.PreOrder(node.RNode)
	println(&node.Data)
}
