package graph

import (
	"container/list"
	"fmt"
)

type BreadthFirstPath struct {
	marked map[int]bool
	edgeTo map[int]int
	startV int // 起点
}

// 广度优先寻路
func NewBreadthFirstPath(sgraph BaseGraph, s int) *BreadthFirstPath {
	b := &BreadthFirstPath{
		marked: make(map[int]bool),
		edgeTo: make(map[int]int),
		startV: 0,
	}

	b.bfs(sgraph, s)

	// fmt.Println("edge to : " ,b.edgeTo)
	return b
}

func (b *BreadthFirstPath) bfs(sgraph BaseGraph, s int) {
	queue := list.New()
	b.marked[s] = true
	queue.PushBack(s)

	for ; queue.Len()>0;  {
		v := queue.Front()
		queue.Remove(v)

		adj,_ := sgraph.Adj(v.Value.(int))
		for _,w := range adj {
			isMarked,ok := b.marked[w.(int)]
			if !ok || !isMarked {
				b.edgeTo[w.(int)] = v.Value.(int)
				b.marked[w.(int)] = true
				queue.PushBack(w.(int))
			}
		}
	}
}

func (b *BreadthFirstPath) hasPathTo(v int) bool {

	isMarked,ok := b.marked[v]
	if !ok {
		return false
	}

	return isMarked
}

func (b *BreadthFirstPath) PathTo(v int) (error,[]int) {

	var path []int
	if !b.hasPathTo(v) {
		return fmt.Errorf("does't have any path to %d", v),path
	}

	for x:=v; x!=b.startV; x=b.edgeTo[x] {

		path = append(path, x)
	}

	path = append(path, b.startV)
	return nil,path
}

