package graph

import "sync"

// 有向加权图
type DirectedWeightEdge struct {
	// 边：这个例子只需要存储一个两个节点的路程，实际应用可以在边中描述更多信息
	Distance int32
	From interface{}
	To interface{}
}

// 新建一个有向加权图
func NewDirectedWeightEdge (distance int32, f , t interface{}) *DirectedWeightEdge {
	return &DirectedWeightEdge{
		distance ,
		f,
		t,
	}
}

type DirectedWeightGraph struct {
	mutex         sync.RWMutex
	adjacencyList map[interface{}]map[interface{}]*DirectedWeightEdge
	v, e          int
}


// V returns the number of vertices in the DirectedGraph
func (g *DirectedWeightGraph) V() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	return g.v
}

// E returns the number of edges in the DirectedGraph
func (g *DirectedWeightGraph) E() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	return g.e
}

// AddEdge will create an edge from vertices v to w
func (g *DirectedWeightGraph) AddEdge(v, w interface{}, distance int32) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if v == w {
		return ErrSelfLoop
	}

	g.addVertex(v)
	g.addVertex(w)

	if _, ok := g.adjacencyList[v][w]; ok {
		return ErrParallelEdge
	}

	g.adjacencyList[v][w] = NewDirectedWeightEdge(distance,v,w)
	g.e++
	return nil
}

// Adj returns the list of all vertices connected to v
func (g *DirectedWeightGraph) Adj(v interface{}) (map[interface{}]*DirectedWeightEdge, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	_, err := g.Degree(v)
	if err != nil {
		return nil, ErrVertexNotFound
	}

	return  g.adjacencyList[v], nil
}

// Degree returns the number of vertices connected to v
func (g *DirectedWeightGraph) Degree(v interface{}) (int, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	val, ok := g.adjacencyList[v]
	if !ok {
		return 0, ErrVertexNotFound
	}
	return len(val), nil
}

func (g *DirectedWeightGraph) addVertex(v interface{}) {
	mm, ok := g.adjacencyList[v]
	if !ok {
		mm = make(map[interface{}]*DirectedWeightEdge)
		g.adjacencyList[v] = mm
		g.v++
	}
}


// 返回所有节点
func (g *DirectedWeightGraph) Vertexes() []interface{} {
	vertexes := make([]interface{}, 0)
	for k,_ := range g.adjacencyList {
		vertexes = append(vertexes, k)
	}

	return vertexes
}

// NewDirectedGraph creates and returns a DirectedGraph
func NewDirectedWeightGraph() *DirectedWeightGraph {
	return &DirectedWeightGraph{
		adjacencyList: make(map[interface{}]map[interface{}]*DirectedWeightEdge),
		v:             0,
		e:             0,
	}
}