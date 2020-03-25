package graph

import (
	"fmt"
	"github.com/phpxin/go-datastruct/utils"
)

// 有向加权图搜索
type DirectedWeightGraphSearch struct {
	Src                 int32           // 根节点，须要搜索的原点
	TotalCostToThisNode map[int32]int32 // 保存原点到每个节点的权
	SPT                 *utils.SPTNode        // 最小生成树根节点
	Graph               *DirectedWeightGraph
}

// 新建有向加权图搜索
func NewDirectedWeightGraphSearch(src int32, g *DirectedWeightGraph) *DirectedWeightGraphSearch {
	return &DirectedWeightGraphSearch{
		Src:                 src,
		TotalCostToThisNode: nil,
		Graph:               g,
	}
}

// 返回 n -> t 的权
func (d *DirectedWeightGraphSearch) EdgeCost(n, t int32) (int32, bool) {
	adjs, _ := d.Graph.Adj(n)

	for k, v := range adjs {
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
	m_CostToThisNode := make(map[int32]int32)
	m_SearchFrontier := make(map[int32]*DirectedWeightEdge)
	m_ShortestPathTree := make(map[int32]*DirectedWeightEdge)

	vertexes := d.Graph.Vertexes()
	for _, v := range vertexes {
		vi := v.(int32)
		m_CostToThisNode[vi] = 0
		m_SearchFrontier[vi] = nil
		m_ShortestPathTree[vi] = nil
	}

	ipq := utils.NewIndexedPriorityQueue(func(a interface{}, b interface{}) int32 {
		a_int := a.(int32)
		b_int := b.(int32)

		if m_CostToThisNode[a_int] == m_CostToThisNode[b_int] {
			return 0
		}

		if m_CostToThisNode[a_int] > m_CostToThisNode[b_int] {
			return 1
		}

		return -1
	})

	ipq.Insert(d.Src)

	for !ipq.Empty() {

		nextClosestNode := ipq.Pop().(int32)

		m_ShortestPathTree[nextClosestNode] = m_SearchFrontier[nextClosestNode]

		adjs, _ := d.Graph.Adj(nextClosestNode)

		for _vertex, edge := range adjs {
			vertex, ok := _vertex.(int32)
			if !ok {
				panic("convert ADJ type failed")
			}

			NewCost := edge.Distance + m_CostToThisNode[nextClosestNode]

			if m_SearchFrontier[vertex] == nil {
				// vertex is pe->To()
				m_CostToThisNode[vertex] = NewCost
				ipq.Insert(vertex)
				m_SearchFrontier[vertex] = edge
			} else if NewCost < m_CostToThisNode[vertex] && m_ShortestPathTree[vertex] == nil {
				m_CostToThisNode[vertex] = NewCost
				ipq.ChangePriority(vertex)
				m_SearchFrontier[vertex] = edge
			}
		}
	}

	// create SPT
	mstArr := make(map[int32]*utils.SPTNode)
	d.SPT = utils.NewSPTNode(d.Src)
	mstArr[d.Src] = d.SPT

	for _, v := range m_ShortestPathTree {
		if v == nil {
			continue
		}

		sptFromNode,ok := mstArr[v.From.(int32)]
		if !ok {
			sptFromNode = utils.NewSPTNode(v.From.(int32))
			mstArr[v.From.(int32)] = sptFromNode
		}

		sptToNode,ok := mstArr[v.To.(int32)]
		if !ok {
			sptToNode = utils.NewSPTNode(v.To.(int32))
			mstArr[v.To.(int32)] = sptToNode
		}

		sptFromNode.AddChild(sptToNode)

		sptToNode.Parent = sptFromNode
		sptToNode.Weight = v.Distance
	}

	d.TotalCostToThisNode = m_CostToThisNode
}
