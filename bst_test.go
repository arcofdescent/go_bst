package bst

import "testing"
import "reflect"
import "math/rand"

func TestNewRoot(t *testing.T) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(5, ch)

		if reflect.TypeOf(rootNode).String() != "*bst.Node" {
			t.Errorf("NewRoot()")
		}
	}()

	<-ch
}

func TestAddNode(t *testing.T) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(5, ch)
		rootNode.AddNode(4)
		rootNode.AddNode(5)

		if rootNode.Search(4) != true {
			t.Errorf("AddNode(): 4 not found")
		}
	}()

	<-ch
}

func TestGetItems(t *testing.T) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(5, ch)
		rootNode.AddNode(4)
		rootNode.AddNode(8)
		rootNode.AddNode(6)
		rootNode.AddNode(2)
		rootNode.AddNode(-2)

		expected := []int{-2, 2, 4, 5, 6, 8}
		items := rootNode.GetItems()
		if !reflect.DeepEqual(items, expected) {
			t.Errorf(`GetItems()`)
		}
	}()

	<-ch
}

func TestSearch(t *testing.T) {

	ch := make(chan int)

	go func() {

		// our test data
		tests := []struct {
			input    int
			expected bool
		}{
			{1, true},
			{3, false},
			{-3, true},
			{6, true},
			{7, true},
			{9, false},
			{89, false},
			{42, true}, // 42 has to be true!
		}

		// lets first add the nodes

		// keep the root node available for later scope
		rootNode := NewRoot(5, ch)

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

	}()

	<-ch
}

func TestDeleteNodeLeaf(t *testing.T) {

	ch := make(chan int)

	go func() {

		rootNode := NewRoot(5, ch)
		rootNode.AddNode(4)

		// delete (left child)
		rootNode.DeleteNode(4)

		if rootNode.Search(4) == true {
			t.Errorf("DeleteNode(): 4 found")
		}

		// delete (right child)
		rootNode.AddNode(6)
		rootNode.DeleteNode(6)

		if rootNode.Search(6) == true {
			t.Errorf("DeleteNode(): 6 found")
		}

	}()

	<-ch
}

func TestDeleteNodeOneChildLeft(t *testing.T) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(5, ch)
		rootNode.AddNode(4)
		rootNode.AddNode(3)

		rootNode.DeleteNode(4)

		if rootNode.Search(4) == true {
			t.Errorf("DeleteNode(): 4 found")
		}

	}()

	<-ch
}

func TestDeleteNodeOneChildRight(t *testing.T) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(5, ch)
		rootNode.AddNode(3)
		rootNode.AddNode(4)

		rootNode.DeleteNode(3)

		if rootNode.Search(3) == true {
			t.Errorf("DeleteNode(): 3 found")
		}
	}()

	<-ch
}

func TestDeleteNodeTwoChild(t *testing.T) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(6, ch)
		rootNode.AddNode(4)
		rootNode.AddNode(3)
		rootNode.AddNode(5)

		rootNode.DeleteNode(4)

		if rootNode.Search(4) == true {
			t.Errorf("DeleteNode(): 4 found")
		}
	}()

	<-ch
}

func TestDuplicate(t *testing.T) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(6, ch)
		rootNode.AddNode(4)
		rootNode.AddNode(4)

		expected := []int{4, 6}
		items := rootNode.GetItems()
		if !reflect.DeepEqual(items, expected) {
			t.Errorf(`Duplicate`)
		}
	}()

	<-ch
}

// test all types of delete
// just ensure no panic takes place :p
func TestDeleteAll(t *testing.T) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(6, ch)

		// generate some random numbers
		nums := []int{}

		for i := 0; i < 50; i++ {
			nums = append(nums, rand.Intn(49))
		}

		// insert
		for _, val := range nums {
			rootNode.AddNode(val)
		}

		// delete
		for _, val := range nums {
			if val != 6 {
				rootNode.DeleteNode(val)
			}
		}
	}()

	<-ch
}
