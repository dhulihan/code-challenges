package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Data rune

func (d Data) String() string {
	return fmt.Sprintf("%c", rune(d))
}

type Node struct {
	data        Data
	left, right *Node
}

func (root *Node) PreOrder(nodes []Data) []Data {
	if root == nil {
		return nodes
	}

	log.Debugf("Performing PreOrder operation on %#v, nodes currently: %s", root.data, nodes)
	nodes = append(nodes, root.data)

	if root.left != nil {
		nodes = root.left.PreOrder(nodes)
	}

	if root.right != nil {
		nodes = root.right.PreOrder(nodes)
	}

	return nodes
}

func (root *Node) InOrder(nodes []Data) []Data {
	if root == nil {
		return nodes
	}

	if root.left != nil {
		nodes = root.left.InOrder(nodes)
	}

	log.Debugf("Performing InOrder operation on %#v, nodes currently: %s", root.data, nodes)
	nodes = append(nodes, root.data)

	if root.right != nil {
		nodes = root.right.InOrder(nodes)
	}

	return nodes
}

func (root *Node) PostOrder(nodes []Data) []Data {
	if root == nil {
		return nodes
	}

	if root.left != nil {
		nodes = root.left.PostOrder(nodes)
	}

	if root.right != nil {
		nodes = root.right.PostOrder(nodes)
	}

	log.Debugf("Performing PostOrder operation on %#v, nodes currently: %s", root.data, nodes)
	nodes = append(nodes, root.data)

	return nodes
}

func (root *Node) LevelOrder(nodes []Data) []Data {
	queue := []*Node{root}

	for len(queue) > 0 {
		var n *Node

		// Pop
		n, queue = queue[0], queue[1:]

		log.Debugf("Performing LevelOrder operation on %#v, nodes currently: %s", root.data, nodes)
		nodes = append(nodes, n.data)

		if n.left != nil {
			queue = append(queue, n.left)
		}

		if n.right != nil {
			queue = append(queue, n.right)
		}
	}

	return nodes
}

func NewLeaf(data Data) *Node {
	return &Node{
		data: data,
	}
}

func NewNode(data Data, left, right *Node) *Node {
	return &Node{
		data:  data,
		left:  left,
		right: right,
	}
}

func main() {
	log.SetLevel(log.DebugLevel)

	// create nodes
	n4 := NewLeaf('E')
	n3 := NewLeaf('D')
	n2 := NewLeaf('C')
	n1 := NewNode('B', n3, n4)
	root := NewNode('A', n1, n2)

	// empty list
	empty := []Data{}

	// do operations
	fmt.Println("Pre Order:", root.PreOrder(empty))
	fmt.Println("In Order:", root.InOrder(empty))
	fmt.Println("Post Order:", root.PostOrder(empty))
	fmt.Println("Level Order:", root.LevelOrder(empty))
}
