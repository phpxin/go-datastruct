package graph

import "sync"

// 有向图，有环

type DirectedGraph struct {
	mutex         sync.RWMutex
	adjacencyList map[interface{}]map[interface{}]struct{}
	v, e          int
}

// V returns the number of vertices in the DirectedGraph
func (g *DirectedGraph) V() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	return g.v
}

// E returns the number of edges in the DirectedGraph
func (g *DirectedGraph) E() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	return g.e
}

// AddEdge will create an edge from vertices v to w
func (g *DirectedGraph) AddEdge(v, w interface{}) error {
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

	g.adjacencyList[v][w] = struct{}{}
	//g.adjacencyList[w][v] = struct{}{}
	g.e++
	return nil
}

// Adj returns the list of all vertices connected to v
func (g *DirectedGraph) Adj(v interface{}) ([]interface{}, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	deg, err := g.Degree(v)
	if err != nil {
		return nil, ErrVertexNotFound
	}

	adj := make([]interface{}, deg)
	i := 0
	for key := range g.adjacencyList[v] {
		adj[i] = key
		i++
	}
	return adj, nil
}

// Degree returns the number of vertices connected to v
func (g *DirectedGraph) Degree(v interface{}) (int, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	val, ok := g.adjacencyList[v]
	if !ok {
		return 0, ErrVertexNotFound
	}
	return len(val), nil
}

func (g *DirectedGraph) addVertex(v interface{}) {
	mm, ok := g.adjacencyList[v]
	if !ok {
		mm = make(map[interface{}]struct{})
		g.adjacencyList[v] = mm
		g.v++
	}
}

// NewDirectedGraph creates and returns a DirectedGraph
func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		adjacencyList: make(map[interface{}]map[interface{}]struct{}),
		v:             0,
		e:             0,
	}
}