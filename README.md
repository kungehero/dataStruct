# dataStruct 

# listzk 链表的实现
```
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
```

# stack 栈的实现
```
func (s *Stack) Push(x interface{}) error {
	Data = make([]interface{}, 0, 10)
	length, capacity := len(Data), cap(Data)
	if length == capacity {
		return errors.New("栈已满！")
	}
	Data = Data[:length+1]
	Data[length] = x
	println(len(Data))
	return nil
}
```

# queue 队列的实现
```
//进队列
func (list *Queue) EnQueue(x interface{}) error {
	list.Add(x)
	return nil
}
```

# binary tree 二叉树的实现

## travel 非遍历实现前中后序二叉树，遍历高度

## reversetree 反转二叉树(镜像树)
```
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
```
