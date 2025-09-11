package linked_lists

func getNode(listHead *SinglyLinkedListNode, positionFromTail int32) int32 {
	mainPointer := listHead
	aheadPointer := listHead

	// two pointers, one pointer i will move ahead on the list positionFromTail times
	var i int32
	for i = 0; i < positionFromTail; i++ {
		aheadPointer = aheadPointer.next
	}

	// move both pointers until aheadPointer reaches the last node,
	// if the aheadPointer is at the last node that means the mainPointer is the node we are looking for (because of original positionFromTail advance)
	for aheadPointer.next != nil {
		mainPointer = mainPointer.next
		aheadPointer = aheadPointer.next
	}

	return mainPointer.data
}
