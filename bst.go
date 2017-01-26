// sorted binary tree
package bst

import "fmt"

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

	// check val does not already exist
	if n.Search(val) {
		return
	}

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

	if !n.Search(val) {
		return
	}

	if val < n.Data {
		n.LeftChild.DeleteNode(val)
	} else if val > n.Data {
		n.RightChild.DeleteNode(val)
	} else {

		if n.LeftChild != nil && n.RightChild != nil {
			successor := n.RightChild.findMinNode()
			n.Data = successor.Data
			successor.DeleteNode(successor.Data)
		} else if n.LeftChild != nil {
			fmt.Println("LeftChild")
			oneChildDelete(n, n.LeftChild)
		} else if n.RightChild != nil {
			fmt.Println("RightChild")
			oneChildDelete(n, n.RightChild)
		} else {
			// leaf node
			parent := n.Parent

			if n.Data < parent.Data {
				parent.LeftChild = nil
			} else {
				parent.RightChild = nil
			}
		}

	}
}

// This function handles the replacement of the node with its child in the case
// of deleting a node with one child
func oneChildDelete(node *Node, childNode *Node) {
	if node.Parent != nil {
		if node.Parent.LeftChild != nil {
			node.Parent.LeftChild = childNode
		} else {
			node.Parent.RightChild = childNode
		}
	}

	childNode = node.Parent
}

func (n *Node) findMinNode() *Node {
	for n.LeftChild != nil {
		n = n.LeftChild
	}

	return n
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
