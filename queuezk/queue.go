package queuezk

import (
	"dataStruct/listzk"
	"errors"
)

type Queue struct {
	listzk.List
}

//进队列
func (list *Queue) EnQueue(x interface{}) error {
	list.Add(x)
	return nil
}

//出队列
func (list *Queue) DeQueue() (x interface{}, err error) {
	if list.GetHeadNode() == nil {
		return "", errors.New("队列为空！")
	} else {
		headnode := list.GetHeadNode()
		list.RemoveAtIndex(0)
		return headnode, nil
	}
}

//获取队列长度
func (list *Queue) GetSize() interface{} {
	return list.Length()
}
