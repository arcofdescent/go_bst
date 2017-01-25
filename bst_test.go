package bst

import "testing"
import "reflect"

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

/*

func TestAddNode(t *testing.T) {
}

func TestSearch(t *testing.T) {
}

func TestDeleteNode(t *testing.T) {
}

func TestImmutability(t *testing.T) {
}

func TestConcurrency(t *testing.T) {
}

func TestMemUsage(t *testing.T) {
}

*/
