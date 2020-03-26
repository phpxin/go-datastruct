package graph

import (
	"fmt"
	"github.com/phpxin/go-datastruct/utils"
)

// 有向加权图搜索
type DirectedWeightGraphSearch struct {
	Src                 int32           // 根节点，须要搜索的原点
	TotalCostToThisNode map[int32]int32 // 保存原点到每个节点的权
	SPT                 *utils.SPTNode  // 最小生成树根节点
	SPTNodeMap map[int32]*utils.SPTNode // 最小生成树全部节点的映射
	Graph               *DirectedWeightGraph
}

// 新建有向加权图搜索
func NewDirectedWeightGraphSearch(src int32, g *DirectedWeightGraph) *DirectedWeightGraphSearch {
	return &DirectedWeightGraphSearch{
		Src:                 src,
		TotalCostToThisNode: nil,
		SPT:                 nil,
		SPTNodeMap: nil,
		Graph:               g,
	}
}

//检索一条通往 target 的路径
func (d *DirectedWeightGraphSearch) Search(target int32) []int32 {
	if d.SPT==nil {
		panic("please create SPT first")
	}

	result := make([]int32, 0)
	result = append(result, target)

	if target==d.Src{
		return result
	}

	targetNode,ok := d.SPTNodeMap[target]
	if !ok {
		panic("target doesn't exist")
	}

	ll := utils.NewLinkedList()
	ll.Push(targetNode)

	for !ll.Empty() {
		item,err := ll.Pop()
		if err!=nil {
			panic(err)
		}

		parent := item.(*utils.SPTNode).Parent.Value

		result = append(result, parent)

		if parent!=d.Src {
			ll.Push(item.(*utils.SPTNode).Parent)
		}

	}

	return result
}

// 返回 n -> t 的权
func (d *DirectedWeightGraphSearch) EdgeCost(n, t int32) (int32, bool) {
	adjCollection, _ := d.Graph.Adj(n)

	for k, v := range adjCollection {
		kk, ok := k.(int32)
		if !ok {
			panic("convert ADJ type failed")
		}

		if kk == t {
			return v.Distance, true
		}
	}

	return 0, false
}

func (d *DirectedWeightGraphSearch) PrintAllPathCosts() {

	for k, v := range d.TotalCostToThisNode {
		fmt.Println(fmt.Sprintf("%d to %d costs %d", d.Src, k, v))
	}

}

// 最小生成树
func (d *DirectedWeightGraphSearch) CreateSPT() {
	mCostToThisNode := make(map[int32]int32)
	mSearchFrontier := make(map[int32]*DirectedWeightEdge)
	mShortestPathTree := make(map[int32]*DirectedWeightEdge)

	vertexes := d.Graph.Vertexes()
	for _, v := range vertexes {
		vi := v.(int32)
		mCostToThisNode[vi] = 0
		mSearchFrontier[vi] = nil
		mShortestPathTree[vi] = nil
	}

	ipq := utils.NewIndexedPriorityQueue(func(a interface{}, b interface{}) int32 {
		aInt := a.(int32)
		bInt := b.(int32)

		if mCostToThisNode[aInt] == mCostToThisNode[bInt] {
			return 0
		}

		if mCostToThisNode[aInt] > mCostToThisNode[bInt] {
			return 1
		}

		return -1
	})

	ipq.Insert(d.Src)

	for !ipq.Empty() {

		nextClosestNode := ipq.Pop().(int32)

		mShortestPathTree[nextClosestNode] = mSearchFrontier[nextClosestNode]

		adjs, _ := d.Graph.Adj(nextClosestNode)

		for _vertex, edge := range adjs {
			vertex, ok := _vertex.(int32)
			if !ok {
				panic("convert ADJ type failed")
			}

			NewCost := edge.Distance + mCostToThisNode[nextClosestNode]

			if mSearchFrontier[vertex] == nil {
				mCostToThisNode[vertex] = NewCost
				ipq.Insert(vertex)
				mSearchFrontier[vertex] = edge
			} else if NewCost < mCostToThisNode[vertex] && mShortestPathTree[vertex] == nil {
				mCostToThisNode[vertex] = NewCost
				ipq.ChangePriority(vertex)
				mSearchFrontier[vertex] = edge
			}
		}
	}

	// create SPT
	d.SPTNodeMap = make(map[int32]*utils.SPTNode)
	d.SPT = utils.NewSPTNode(d.Src)
	d.SPTNodeMap[d.Src] = d.SPT

	for _, v := range mShortestPathTree {
		if v == nil {
			continue
		}

		sptFromNode, ok := d.SPTNodeMap[v.From.(int32)]
		if !ok {
			sptFromNode = utils.NewSPTNode(v.From.(int32))
			d.SPTNodeMap[v.From.(int32)] = sptFromNode
		}

		sptToNode, ok := d.SPTNodeMap[v.To.(int32)]
		if !ok {
			sptToNode = utils.NewSPTNode(v.To.(int32))
			d.SPTNodeMap[v.To.(int32)] = sptToNode
		}

		sptFromNode.AddChild(sptToNode)

		sptToNode.Parent = sptFromNode
		sptToNode.Weight = v.Distance
	}

	d.TotalCostToThisNode = mCostToThisNode
}
