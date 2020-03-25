package utils

import "fmt"

// 优先队列,一个从小到大排列的数组
// 注意：该优先队列查找和插入算法在大数据集情况下都比较低效
// 这里只是简单地学习使用，生产环境请适度改变排序算法，推荐根据数据量选择线性排序，二分排序，希尔排序
type IndexedPriorityQueue struct {
	queue []interface{}
	compareFunc func (interface{}, interface{}) int32 // a > b returns 1, a = b returns 0, a < b returns -1

}

// 新建优先队列
func NewIndexedPriorityQueue( cf func (interface{}, interface{}) int32 ) *IndexedPriorityQueue {
	return &IndexedPriorityQueue{
		queue: make([]interface{}, 0) ,
		compareFunc: cf ,
	}
}

// 打印队列元素
func (q *IndexedPriorityQueue) Test() {
	fmt.Println(q.queue)
}

// 队列是否为空
func (q *IndexedPriorityQueue) Empty() bool {
	if len(q.queue)<=0 {
		return true
	}

	return false
}

// 重排成员优先级
func (q *IndexedPriorityQueue) ChangePriority(item interface{}) {
	// delete it first

	newQueue := make([]interface{}, 0)

	lenOfQueue := len(q.queue)

	for index:=0; index<lenOfQueue; index++ {
		compRet := q.compareFunc(item, q.queue[index])
		if compRet!=0 {
			newQueue = append(newQueue, q.queue[index])
		}
	}

	q.queue = newQueue

	q.Insert(item) // then insert it
}

// 插入成员
func (q *IndexedPriorityQueue) Insert(item interface{}){
	lenOfQueue := len(q.queue)
	index:=0
	for index=0; index<lenOfQueue; index++ {
		compRet := q.compareFunc(item, q.queue[index])
		if compRet<=0 {
			break
		}
	}

	if index==lenOfQueue {
		q.queue = append(q.queue, item)
	}else{
		q.queue = append(q.queue, struct{}{})
		for i:=lenOfQueue; i>index; i-- {
			q.queue[i] = q.queue[i-1]
		}
	}

	q.queue[index] = item
}

// 弹出最优先成员
func (q *IndexedPriorityQueue) Pop() interface{} {

	if len(q.queue)==0 {
		return nil
	}

	ret := q.queue[0]

	q.queue = q.queue[1:len(q.queue)]

	return ret
}