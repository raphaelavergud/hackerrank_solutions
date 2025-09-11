package binary_trees

import (
	"fmt"
)

//	order of traversal:
//
// left -> right -> root
func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%d ", root.data)
}
