package bst

import "testing"
import "reflect"

//import "fmt"

/*
	tests := []struct {
		input int
		expected int
	}{
		{1, 4},
	}
*/

func TestNewRoot(t *testing.T) {
	rootNode := NewRoot(5)

	if reflect.TypeOf(rootNode).String() != "*bst.Node" {
		t.Errorf("NewRoot()")
	}
}

func TestAddNode(t *testing.T) {
	rootNode := NewRoot(5)
	rootNode.AddNode(4)

	if rootNode.Search(4) != true {
		t.Errorf("AddNode(): 4 not found")
	}
}

func TestGetItems(t *testing.T) {
	rootNode := NewRoot(5)
	rootNode.AddNode(4)
	rootNode.AddNode(8)
	rootNode.AddNode(6)
	rootNode.AddNode(2)

	expected := []int{2, 4, 5, 6, 8}
	items := rootNode.GetItems()
	if !reflect.DeepEqual(items, expected) {
		t.Errorf(`GetItems()`)
	}
}

func TestSearch(t *testing.T) {
	// our test data
	tests := []struct {
		input    int
		expected bool
	}{
		{1, true},
		{3, false},
		{6, true},
		{7, true},
		{9, false},
		{89, false},
		{42, true}, // 42 has to be true!
	}

	// lets first add the nodes

	// keep the root node available for later scope
	rootNode := NewRoot(5)

	// children
	for _, tt := range tests {
		if tt.expected {
			rootNode.AddNode(tt.input)
		}
	}

	// OK, now search and verify

	// root has to be true
	if rootNode.Search(5) != true {
		t.Errorf("Search() root node not found")
	}

	// children
	for _, tt := range tests {
		if tt.expected {
			if rootNode.Search(tt.input) != true {
				t.Errorf("Search() node %d not found", tt.input)
			}
			if rootNode.Search(tt.input) == false {
				t.Errorf("Search() node %d found in tree", tt.input)
			}
		}
	}
}

func TestDeleteNode(t *testing.T) {
	rootNode := NewRoot(5)
	rootNode.AddNode(4)
	rootNode.AddNode(3)

	// delete
	rootNode.DeleteNode(4)

	if rootNode.Search(4) == true {
		t.Errorf("DeleteNode(): 4 found")
	}
}

/*
func TestImmutability(t *testing.T) {
}

func TestConcurrency(t *testing.T) {
}

func TestMemUsage(t *testing.T) {
}

func TestSanity(t *testing.T) {
}

func TestDuplicate(t *testing.T) {
}
*/
