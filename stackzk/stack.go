package stackzk

import (
	"errors"
)

var Data []interface{}

type Stack struct {
}

func (s *Stack) Push(x interface{}) error {
	Data = make([]interface{}, 0, 100000)
	length, capacity := len(Data), cap(Data)
	if length == capacity {
		return errors.New("栈已满！")
	}
	Data = Data[:length+1]
	Data[length] = x
	println(len(Data))
	return nil
}

func (s *Stack) Pop() (x interface{}, err error) {
	Data = make([]interface{}, 0, 100000)
	lengtn := len(Data)
	if lengtn == 0 {
		return "", errors.New("栈为空！")
	}
	Data[lengtn-1] = x
	Data = Data[:lengtn-1]
	return x, nil
}
