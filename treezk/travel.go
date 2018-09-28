package treezk

import (
	"container/list"
)

//非递归二叉树遍历

//先序遍历
func (tree *Node) PreOrder() []interface{} {
	array := make([]interface{}, 0)
	stack := list.New()
	for tree != nil || stack.Len() != 0 {
		for tree != nil {
			stack.PushBack(tree)
			array = append(array, tree.Data)
			tree = tree.LNode
		}
		if stack.Len() != 0 {
			v := stack.Back()
			tree = v.Value.(*Node)
			tree = tree.RNode
			stack.Remove(v)
		}
	}
	return array
}

//中序遍历
func (tree *Node) InOrder() []interface{} {
	array := make([]interface{}, 0)
	stack := list.New()
	for tree != nil || stack.Len() != 0 {
		for tree != nil {
			stack.PushBack(tree)
			tree = tree.LNode
		}
		if stack.Len() != 0 {
			array = append(array, tree.Data)
			v := stack.Back()
			tree = v.Value.(*Node)
			tree = tree.RNode
			stack.Remove(v)
		}
	}
	return array
}

//后序遍历

func (tree *Node) BackOrder() []interface{} {
	array := make([]interface{}, 0)
	stack := list.New()
	for tree != nil || stack.Len() != 0 {
		for tree != nil {
			stack.PushBack(tree)
			tree = tree.LNode
		}
		v := stack.Back()
		tree = v.Value.(*Node)
		if tree.LNode == nil && tree.RNode == nil {
			array = append(array, tree.Data)
			stack.Remove(v)
		} else {
			tree = tree.RNode
		}

	}
	return array
}
