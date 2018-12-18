package main

import (
	"awesomeProject/lang/tree/tree"
	"fmt"
)

type mytreeNode struct {
	node *tree.Node
}

func (mynode *mytreeNode) postOrder()  {
	if mynode == nil || mynode.node == nil{
		return
	}
	left:=mytreeNode{mynode.node.Left}
	right:=mytreeNode{mynode.node.Right}
	left.postOrder()
	right.postOrder()
	mynode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	//root.right.left.print()
	//
	//root.print()

	root.Traverse()
	//fmt.Println()
	//mynode := mytreeNode{&root}
	//mynode.postOrder()

	/*var proot *treenode
	proot.setValue(200)
	proot = &root
	proot.setValue(300)
	proot.print()*/

	countsTraver := 0
	root.TraverseFunc(func(node *tree.Node) {
		countsTraver++
	})
	fmt.Println("node count" ,countsTraver)
}
