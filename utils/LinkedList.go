package utils

import "fmt"

// 链表
type LinkedList struct {
	data []interface{}
	len int
}

// 新建链表
func NewLinkedList() *LinkedList {
	return &LinkedList{
		data:make([]interface{},0),
		len:0,
	}
}

// 链表长度
func (l *LinkedList) Len() int32 {
	return int32(l.len)
}

// 是否为空链表
func (l *LinkedList) Empty() bool {
	if l.len<=0 {
		return true
	}

	return false
}

// 从末尾键入一条数据
func (l *LinkedList) Push(item interface{}) {
	l.data = append(l.data, item)
	l.len = l.len+1
}

// 从表头弹出一条数据（队列方法）
func (l *LinkedList) Pop() (item interface{}, err error) {
	l.len = len(l.data)
	if l.len > 0 {
		item = l.data[0]
		l.len = l.len-1
		if l.len<=0 {
			l.data = make([]interface{}, 0) // flush
		}else{
			l.data = l.data[1:]
		}

		return item,nil
	}

	return nil, fmt.Errorf("no enough items")
}

// 从末尾弹出一条数据（栈方法）
func (l *LinkedList) Shift() (item interface{}, err error) {
	l.len = len(l.data)
	if l.len > 0 {
		item = l.data[l.len-1]
		l.len = l.len-1
		if l.len<=0 {
			l.data = make([]interface{}, 0) // flush
		}else{
			l.data = l.data[:l.len]
		}

		return item,nil
	}

	return nil, fmt.Errorf("no enough items")
}