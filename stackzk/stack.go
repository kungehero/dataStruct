package stackzk

import (
	"errors"
)

var data []interface{}

func Push(x interface{}) error {
	data = make([]interface{}, 0, 100000)
	length, capacity := len(data), cap(data)
	if length == capacity {
		return errors.New("栈已满！")
	}
	data = data[:length]
	data[length-1] = x
	return nil
}

func Pop() (x interface{}, err error) {
	data = make([]interface{}, 0, 100000)
	lengtn := len(data)
	if lengtn == 0 {
		return "", errors.New("栈为空！")
	}
	data[lengtn-1] = x
	data = data[:lengtn-1]
	return x, nil
}
