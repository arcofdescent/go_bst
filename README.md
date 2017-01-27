
###Binary Search Tree implemented in Go (golang)

* First download Go from <https://golang.org/dl/> and install it
* Clone this repo

Then...

```
cd nilenso_bst
```

Run the tests

```
go test
```

Additionally


```
go test -v
```

To see test coverage

```
go test -cover
```

To run a simple benchmark for **AddNode()**

```
go test -bench=.
```

####Usage

```go
package main

import "fmt"
import "nilenso_bst" // assuming you have it installed properly

rootNode := nilenso_bst.NewRoot(7)
rootNode.AddNode(5)
rootNode.AddNode(9)
rootNode.AddNode(8)

if rootNode.Search(5) == true {
  fmt.Println("5 found")
}

if rootNode.Search(10) == false {
  fmt.Println("10 does not exist")
}

rootNode.DeleteNode(9)

if rootNode.Search(9) == false {
  fmt.Println("9 does not exists")
}

items := rootNode.GetItems() // in sorted order
fmt.Printf("%v\n", items) // [5, 7, 8]
```
