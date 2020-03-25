package graph

// 基础图接口
type BaseGraph interface {
	V() int
	E() int
	AddEdge(v, w interface{}) error
	Adj(v interface{}) ([]interface{}, error)
	Degree(v interface{}) (int, error)
	addVertex(v interface{})
	
}