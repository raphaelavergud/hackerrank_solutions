package binary_trees

import (
	"fmt"
)

type Node struct {
	data  int32
	left  *Node
	right *Node
}

//	order of traversal:
//
// root -> left -> right
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}
