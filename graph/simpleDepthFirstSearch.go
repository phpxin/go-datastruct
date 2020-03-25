package graph


type DepthFirstSearch struct {
	marked map[int]bool
	count int
}

// 深度优先遍历
func NewDepthFirstSearch(sgraph *SimpleGraph, s int) *DepthFirstSearch {
	d := &DepthFirstSearch{
		marked: make(map[int]bool),
		count: 0,
	}

	d.dfs(sgraph, s)
	return d
}

func (d *DepthFirstSearch) dfs(sgraph *SimpleGraph, v int) {
	d.marked[v] = true
	d.count++

	adj,_ := sgraph.Adj(v)
	for _,w := range adj {
		isMarked,ok := d.marked[w.(int)]
		if !ok || !isMarked {
			d.dfs(sgraph, w.(int))
		}
	}
}

func (d *DepthFirstSearch) MarkedList() []int {
	var mList []int
	for k,_ := range d.marked {
		mList = append(mList, k)
	}
	return mList
}

func (d *DepthFirstSearch) Marked(w int) bool {

	isMarked,ok := d.marked[w]
	if !ok {
		return false
	}

	return isMarked
}

func (d *DepthFirstSearch) Count() int {
	return d.count
}

