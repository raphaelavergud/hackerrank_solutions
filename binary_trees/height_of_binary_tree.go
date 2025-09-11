package binary_trees

func getHeight(root *Node) int32 {
	// today i learned: by convention, the height of an empty tree is minus 1 (-1)
	// this is also necessary for the following repeats when we call getHeight for the children
	if root == nil {
		return -1
	}

	// recursively find out heights of both sides
	leftSubtreeHeight := getHeight(root.left)
	rightSubtreeHeight := getHeight(root.right)

	// whichever is the taller height is the total height
	var tallerSubtreeHeight int32
	if leftSubtreeHeight > rightSubtreeHeight {
		tallerSubtreeHeight = leftSubtreeHeight
	} else {
		tallerSubtreeHeight = rightSubtreeHeight
	}
	
	return 1 + tallerSubtreeHeight
}
