package main

import (
	"fmt"
	"github.com/phpxin/go-datastruct/graph"
)

// 从矩阵中提取节点，返回编号数组
// 节点编号：数组顺序从左到右，从上到下，从0开始正序排列
func getVertexList(data [][]int)  []int {
	blocks := []int{} // 路块总数
	height := len(data)
	for i:=0;i<height;i++ {
		width := len(data[i])
		for j := 0; j < width; j++ {
			if data[i][j]==1 {
				vertex := i*width+j // 当前节点编号，数组顺序从左到右，从上到下，从0开始正序排列
				blocks = append(blocks, vertex)
			}
		}
	}
	return blocks
}

// 通过一个矩阵生成一副简单图，1 代表图上的一个点
// 节点编号：数组顺序从左到右，从上到下，从0开始正序排列
func generateSimpleGraph(data [][]int) *graph.SimpleGraph {

	sgraph := graph.NewSimpleGraph()
	height := len(data)

	for i:=0;i<height;i++ {
		width := len(data[i])
		for j:=0;j<width;j++ {
			vertex := i*width+j // 当前节点编号，数组顺序从左到右，从上到下，从0开始正序排列

			//纵向节点
			if i<height-1 && data[i][j]==1 && data[i+1][j]==1 {
				// 存在一条边
				vertex2 := (i+1)*width+j
				sgraph.AddEdge(vertex, vertex2)
			}

			// 横向节点
			if data[i][j]==1 && j<width-1 && data[i][j+1]==1 {
				// 存在一条边
				vertex2 := i*width+j+1
				sgraph.AddEdge(vertex, vertex2)
			}
		}
	}

	return sgraph
}


func wayTest(){
	// 路径连通性 Demo

	data := [][]int{
		{1,2,2,1,1},
		{1,1,1,2,1},
		{1,2,3,2,1},
		{1,2,2,2,1},
		{1,1,1,1,1},
	} // 节点总数 16  图节点总数  16  边总数  15  深度优先遍历连接数  16 连通

	//data := [][]int{
	//	{1,2,2,1,1},
	//	{1,1,1,2,2},
	//	{1,2,3,2,1},
	//	{1,2,2,2,1},
	//	{1,1,1,1,1},
	//} // 节点总数 15  图节点总数  15  边总数  13  深度优先遍历连接数  13  不连通

	blocks := getVertexList(data)
	sgraph := generateSimpleGraph(data)
	depthFS := graph.NewDepthFirstSearch(sgraph, blocks[0])

	fmt.Println( "节点总数", len(blocks) ," 图节点总数 ", sgraph.V(), " 边总数 ", sgraph.E(), " 深度优先遍历连接数 ", depthFS.Count())
}
