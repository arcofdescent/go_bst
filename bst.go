// sorted binary tree
package bst

import "sync"

type Node struct {
	data                  int
	leftChild, rightChild *Node
	parent                *Node
	ch                    chan<- int
}

var mu sync.Mutex

func NewRoot(val int, ch chan<- int) *Node {

	defer func() { ch <- 1 }()

	root := Node{}
	root.data = val
	root.ch = ch
	return &root
}

func (n *Node) AddNode(val int) {

	defer func() { n.ch <- 1 }()

	// check that val does not already exist, i,e. no dups
	if n.Search(val) {
		return
	}

	if val < n.data {
		if n.leftChild == nil {
			newNode := &Node{data: val}
			n.leftChild = newNode
			newNode.parent = n
		} else {
			n.leftChild.AddNode(val)
		}
	} else {
		if n.rightChild == nil {
			newNode := &Node{data: val}
			n.rightChild = newNode
			newNode.parent = n
		} else {
			n.rightChild.AddNode(val)
		}
	}
}

func (n *Node) Search(val int) bool {

	defer func() { n.ch <- 1 }()

	// iteratively to prevent a stack overflow
	// normally this is done using recursion which can
	// cause a SO for huge trees
	currNode := n

	for currNode != nil {
		if val == currNode.data {
			return true
		} else if val < currNode.data {
			currNode = currNode.leftChild
		} else {
			currNode = currNode.rightChild
		}
	}

	return false
}

func (n *Node) SearchNode(val int) *Node {

	defer func() { n.ch <- 1 }()

	// iteratively to prevent stack overflow
	currNode := n

	for currNode != nil {
		if val == currNode.data {
			return currNode
		} else if val < currNode.data {
			currNode = currNode.leftChild
		} else {
			currNode = currNode.rightChild
		}
	}

	return nil
}

func (n *Node) DeleteNode(val int) {

	defer func() { n.ch <- 1 }()

	if !n.Search(val) {
		return
	}

	if val < n.data {
		n.leftChild.DeleteNode(val)
	} else if val > n.data {
		n.rightChild.DeleteNode(val)
	} else {

		if n.leftChild != nil && n.rightChild != nil {
			successor := n.rightChild.findMinNode()
			n.data = successor.data
			successor.DeleteNode(successor.data)
		} else if n.leftChild != nil {
			oneChildDelete(n, n.leftChild)
		} else if n.rightChild != nil {
			oneChildDelete(n, n.rightChild)
		} else {
			// leaf node
			parent := n.parent

			if n.data < parent.data {
				parent.leftChild = nil
			} else {
				parent.rightChild = nil
			}
		}

	}
}

// This function handles the replacement of the node
// with its child in the case of deleting a node with one child
func oneChildDelete(node *Node, childNode *Node) {
	if node.parent != nil {
		if node.parent.leftChild != nil {
			node.parent.leftChild = childNode
		} else {
			node.parent.rightChild = childNode
		}
	}

	childNode = node.parent
}

func (n *Node) findMinNode() *Node {
	for n.leftChild != nil {
		n = n.leftChild
	}

	return n
}

func (n *Node) GetItems() []int {

	defer func() { n.ch <- 1 }()

	var vals []int

	// inorder traversal

	// Process leftChild
	if n.leftChild != nil {
		vals = append(vals, n.leftChild.GetItems()...)
	}

	// Process Node
	vals = append(vals, n.data)

	// Process rightChild
	if n.rightChild != nil {
		vals = append(vals, n.rightChild.GetItems()...)
	}

	return vals
}
