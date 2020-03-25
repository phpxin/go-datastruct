package main

import (
	"fmt"
	"github.com/phpxin/go-datastruct/graph"
)

func pathTest(){
	fmt.Println("深度优先寻路:")
	depthFirstPathTest()

	fmt.Println("广度优先寻路:")
	breadthFirstPathTest()
}

func breadthFirstPathTest(){
	// 广度优先寻路
	data := [][]int{
		{1,2,2,1,1},
		{1,1,1,2,1},
		{1,2,1,2,1},
		{1,2,1,2,1},
		{1,1,1,1,1},
	}

	blocks := getVertexList(data)

	startAt := blocks[0]
	endAt := blocks[5]

	fmt.Println("start at ", startAt, " end at ", endAt)

	sgraph := generateSimpleGraph(data)
	bfp := graph.NewBreadthFirstPath(sgraph, startAt)
	err,path := bfp.PathTo(endAt)
	if err!=nil {
		fmt.Println(err)
		return
	}

	for i:=len(path);i>0;i-- {
		fmt.Println(path[i-1])
	}
}

func depthFirstPathTest(){
	// 寻路
	data := [][]int{
		{1,2,2,1,1},
		{1,1,1,2,1},
		{1,2,1,2,1},
		{1,2,1,2,1},
		{1,1,1,1,1},
	}

	blocks := getVertexList(data)

	startAt := blocks[0]
	endAt := blocks[5]

	fmt.Println("start at ", startAt, " end at ", endAt)

	sgraph := generateSimpleGraph(data)
	dfp := graph.NewDepthFirstPath(sgraph, startAt)
	err,path := dfp.PathTo(endAt)
	if err!=nil {
		fmt.Println(err)
		return
	}

	for i:=len(path);i>0;i-- {
		fmt.Println(path[i-1])
	}
}
