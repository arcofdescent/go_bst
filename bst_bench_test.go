package bst

import "testing"

func BenchmarkAddNode(b *testing.B) {

	ch := make(chan int)

	go func() {
		rootNode := NewRoot(5, ch)

		for i := 0; i < 10000; i++ {
			rootNode.AddNode(i)
		}
	}()

	<-ch
}
