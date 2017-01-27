package bst

import "testing"

func BenchmarkAddNode(b *testing.B) {
	rootNode := NewRoot(5)

	for i := 0; i < 1000; i++ {
		rootNode.AddNode(i)
	}
}
