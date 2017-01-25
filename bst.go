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

	// preorder traversal
	// mind the recursion!

	// Process Node
	if n.Data == val {
		return true
	}

	// Process LeftChild
	if n.LeftChild != nil {
		return n.LeftChild.Search(val)
	}

	// Process RightChild
	if n.RightChild != nil {
		return n.RightChild.Search(val)
	}

	return false
}
