package binary_trees

import (
	"fmt"
)

//	order of traversal:
//
// left -> root -> right
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Printf("%d ", root.data)
	inOrder(root.right)
}
