package graph

import "fmt"

type DepthFirstPath struct {
	marked map[int]bool
	edgeTo map[int]int
	startV int // 起点
}

// 深度优先寻路
func NewDepthFirstPath(sgraph *SimpleGraph, s int) *DepthFirstPath {
	d := &DepthFirstPath{
		marked: make(map[int]bool),
		edgeTo: make(map[int]int),
		startV: 0,
	}

	d.dfs(sgraph, s)

	// fmt.Println("edge to : " ,d.edgeTo)
	return d
}

func (d *DepthFirstPath) dfs(sgraph *SimpleGraph, v int) {
	d.marked[v] = true

	adj,_ := sgraph.Adj(v)
	for _,w := range adj {
		isMarked,ok := d.marked[w.(int)]
		if !ok || !isMarked {
			d.edgeTo[w.(int)] = v
			d.dfs(sgraph, w.(int))
		}
	}
}

func (d *DepthFirstPath) hasPathTo(v int) bool {

	isMarked,ok := d.marked[v]
	if !ok {
		return false
	}

	return isMarked
}

func (d *DepthFirstPath) PathTo(v int) (error,[]int) {

	var path []int
	if !d.hasPathTo(v) {
		return fmt.Errorf("does't have any path to %d", v),path
	}

	for x:=v; x!=d.startV; x=d.edgeTo[x] {

		path = append(path, x)
	}

	path = append(path, d.startV)
	return nil,path
}

