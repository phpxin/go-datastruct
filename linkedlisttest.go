package main

import (
	"fmt"
	"github.com/phpxin/go-datastruct/utils"
)

func linkedlist_test() {
	ln := utils.NewLinkedList()
	ln.Push(int32(1))
	ln.Push(int32(2))
	ln.Push(int32(3))
	ln.Push(int32(4))
	ln.Push(int32(5))
	ln.Push(int32(6))
	ln.Push(int32(7))


	for !ln.Empty() {
		item,err := ln.Shift()
		if err!=nil{
			panic(err)
		}

		fmt.Println(item.(int32))
	}
}