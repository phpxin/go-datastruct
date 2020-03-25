package main

import (
	"fmt"
	"github.com/phpxin/go-datastruct/graph"
	"github.com/phpxin/go-datastruct/utils"
	)


// 有向加权图测试
func DWGSearchTest(){

	graphMapSrc := map[int32]map[int32]int32{
		1:{
			2:1,
			4:4,
			5:2,
		},
		2:{
			3:1,
		},
		3:{
			4:1,
			5:1,
		},
	}

	dwg := graph.NewDirectedWeightGraph()

	for k,v := range graphMapSrc {
		for k2,v2 := range v {
			err := dwg.AddEdge(k,k2,v2)
			if err!=nil {
				panic(err)
			}
		}
	}

	dwgs := graph.NewDirectedWeightGraphSearch(1, dwg)
	dwgs.CreateSPT()

	utils.PrintSPT(dwgs.SPT)
	dwgs.PrintAllPathCosts()
}

// 优先队列测试
func IPQTest(){
	compfunc := func (a interface{}, b interface{}) int32 {
		a_int := a.(int32)
		b_int := b.(int32)

		if a_int==b_int {
			return 0
		}

		if a_int>b_int {
			return 1
		}


		return -1
	}

	testarr := []int32{7,3,5,2,9,2,4}

	ipq := utils.NewIndexedPriorityQueue(compfunc)

	for _,v := range testarr {
		ipq.Insert(v)
	}

	for ; ;  {
		//ipq.Test()
		item := ipq.Pop()
		if item==nil {
			break
		}

		fmt.Println(item)
	}
}