package utils

import "fmt"

// 最小生成树节点
type SPTNode struct {
	Parent   *SPTNode
	Value    int32
	Children []*SPTNode
	Weight   int32
}

func NewSPTNode(v int32) *SPTNode {
	return &SPTNode{
		Parent:   nil,
		Value:    v,
		Children: make([]*SPTNode, 0),
		Weight:   0,
	}
}

func PrintSPT(n *SPTNode) {
	from := n.Value
	if n.Parent!=nil {
		from = n.Parent.Value
	}

	fmt.Println(fmt.Sprintf("%d to %d weight is %d", from, n.Value, n.Weight))

	cl := len(n.Children)
	if cl>0 {
		for _,c := range n.Children {
			PrintSPT(c)
		}
	}
}

func (n *SPTNode) AddChild(node *SPTNode) {
	n.Children = append(n.Children, node)
}