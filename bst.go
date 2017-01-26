// sorted binary tree
package bst

//import "fmt"

type Node struct {
	Data                  int
	LeftChild, RightChild *Node
	Parent                *Node
}

func NewRoot(val int) *Node {
	root := Node{}
	root.Data = val
	return &root
}

func (n *Node) AddNode(val int) {
	if val < n.Data {
		if n.LeftChild == nil {
			newNode := &Node{Data: val}
			n.LeftChild = newNode
			newNode.Parent = n
		} else {
			n.LeftChild.AddNode(val)
		}
	} else {
		if n.RightChild == nil {
			newNode := &Node{Data: val}
			n.RightChild = newNode
			newNode.Parent = n
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

func (n *Node) SearchNode(val int) *Node {

	// iteratively to prevent stack overflow
	currNode := n

	for currNode != nil {
		if val == currNode.Data {
			return currNode
		} else if val < currNode.Data {
			currNode = currNode.LeftChild
		} else {
			currNode = currNode.RightChild
		}
	}

	return nil
}

func (n *Node) DeleteNode(val int) {

	// get the node
	node := n.SearchNode(val)
	if node == nil {
		return
	}

	// is the node a leaf i.e. no siblings
	if node.LeftChild == nil && node.RightChild == nil {
		// get rid of it
		parent := node.Parent

		if node.Data < parent.Data {
			parent.LeftChild = nil
		} else {
			parent.RightChild = nil
		}
	}
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
