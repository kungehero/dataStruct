package treezk

import "container/list"

func ReverseTree(node *Node) {
	if node == nil {
		return
	}
	stack := list.New()
	stack.PushFront(node)

	for stack.Len() > 0 {
		countNode := stack.Len()
		for countNode > 0 {
			movenode := stack.Front()
			nextnode := movenode.Value.(*Node)
			if nextnode.LNode != nil {
				stack.PushBack(nextnode.RNode)
			}
			if nextnode.RNode != nil {
				stack.PushBack(nextnode.RNode)
			}

			if nextnode.LNode != nil && nextnode.RNode != nil {
				nextnode.LNode, nextnode.RNode = nextnode.RNode, nextnode.LNode
			}
			if nextnode.LNode != nil && nextnode.RNode == nil {
				nextnode.RNode = nextnode.LNode
			}

			if nextnode.LNode == nil && nextnode.RNode != nil {
				nextnode.LNode = nextnode.RNode
			}
			stack.Remove(movenode)
			countNode--
		}
	}
}
