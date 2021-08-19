package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func initTree() *node {
	root := &node{data: 1,
		left:  nil,
		right: nil}

	return root
}

func tree() {
	root := initTree()
	fmt.Println(root.data)
}
