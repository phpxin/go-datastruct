package graph

type CC struct {
	vertexes []int // 点集
	marked map[int]bool
	id map[int]int
	count int
}

// 基于深度优先遍历计算连通分量
func NewConnectedComponent(vertexes []int, sgraph *SimpleGraph) *CC {
	cc := &CC{
		marked: make(map[int]bool),
		id: make(map[int]int),
		count: 0,
	}

	for _,s := range vertexes {
		isMarked,ok := cc.marked[s]
		if !ok || !isMarked {
			cc.dfs(sgraph, s)
			cc.count++
		}
	}

	return cc
}

func (cc *CC) dfs(sgraph *SimpleGraph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count

	adj,_ := sgraph.Adj(v)
	for _,w := range adj {
		isMarked,ok := cc.marked[w.(int)]
		if !ok || !isMarked {
			cc.dfs(sgraph, w.(int))
		}
	}
}

func (cc *CC) Connected( v int,  w int) bool {
	return cc.id[v] == cc.id[w]
}

func (cc *CC) Id(v int) (int,bool) {
	id,ok := cc.id[v]
	return id,ok
}

func (d *CC) Count() int {
	return d.count
}
