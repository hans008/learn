package main

import "fmt"

type treeNode struct {
	value int
	left,right *treeNode
}


func (node treeNode) Print(){
	fmt.Printf("%v\n",node.value)
}

func (node *treeNode) SetValue(value int){
	node.value = value
}

func (node *treeNode) traverse(){
	//中序遍历
	if node == nil {
		return
	}

	node.left.traverse()
	node.Print()
	node.right.traverse()
}

func createTreeNode(value int) *treeNode {
	return &treeNode{value:value}
}

func main(){
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
	root.left.right = createTreeNode(2)

	root.left.SetValue(10)
	root.left.Print()
	root.Print()

	root.traverse()
}