// sorted binary tree
package bst

//import "fmt"

type Node struct {
	Data                  int
	LeftChild, RightChild *Node
}

func NewRoot(val int) *Node {
	root := Node{}
	root.Data = val
	return &root
}

func (n *Node) AddNode(val int) {
	if val < n.Data {
		if n.LeftChild == nil {
			n.LeftChild = &Node{Data: val}
		} else {
			n.LeftChild.AddNode(val)
		}
	} else {
		if n.RightChild == nil {
			n.RightChild = &Node{Data: val}
		} else {
			n.RightChild.AddNode(val)
		}
	}
}

func (n *Node) Search(val int) bool {

	// iteratively to prevent stack overflow
	currNode := n

	for currNode != nil {
		if val == currNode.Data {
			return true
		} else if val < currNode.Data {
			currNode = currNode.LeftChild
		} else {
			currNode = currNode.RightChild
		}
	}

	return false
}

func (n *Node) DeleteNode(val int) {
}

func (n *Node) GetItems() []int {

	var vals []int

	// inorder traversal

	// Process LeftChild
	if n.LeftChild != nil {
		vals = append(vals, n.LeftChild.GetItems()...)
	}

	// Process Node
	vals = append(vals, n.Data)

	// Process RightChild
	if n.RightChild != nil {
		vals = append(vals, n.RightChild.GetItems()...)
	}

	return vals
}
