package main

import (
	"fmt"
	"github.com/phpxin/go-datastruct/graph"
)

func dg_bfs_test() {
	// 有向图，广度优先遍历，测试

	// 元数据，原点-> 目标点集
	src := map[int][]int{
		1:{2,3} ,
		2:{4},
		3:{1,5},
		4:{2,3},
	}

	var err error

	dg := graph.NewDirectedGraph()

	for srcPoint,dstPoints := range src {
		for _,v := range dstPoints {
			err = dg.AddEdge(srcPoint, v)
			if err!=nil {
				panic(err)
			}
		}
	}

	dfs := graph.NewBreadthFirstPath(dg, 2)
	err,result := dfs.PathTo(5)

	for _,v := range result {
		fmt.Println(v)
	}
}