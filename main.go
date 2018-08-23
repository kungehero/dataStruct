package main

import (
	"dataStruct/treezk"
	"fmt"
)

func main() {
	tree := treezk.Tree{}
	tree.Add(0)
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)

	//广度优先遍历
	tree.BreadthTravel()

	//深度优先 先序遍历
	tree.PreOrder(tree.RootNode)

	//深度优先  中序遍历
	tree.InOrder(tree.RootNode)

	//深度优先  后序遍历
	tree.PostOrder(tree.RootNode)
	fmt.Println("数据结构！")
}
