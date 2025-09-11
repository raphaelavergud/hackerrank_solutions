package linked_lists

func insertNodeAtPosition(listHead *SinglyLinkedListNode, dataToInsert int32, position int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{
		data: dataToInsert,
		next: nil,
	}

	nodeBeforeInsertion := listHead

	var currentPosition int32

	// traverse until we are at 'before' position
	for currentPosition = 0; currentPosition < position-1; currentPosition++ {
		// remember node of the 'before' position, so i can first link newnode.next to nodebefore.next, and then link nodebefore.next to the newnode.
		nodeBeforeInsertion = nodeBeforeInsertion.next
	}

	// has to be in this order,
	// link the new node in between the 'beforenode' and 'neforenode.next'
	newNode.next = nodeBeforeInsertion.next
	nodeBeforeInsertion.next = newNode

	return listHead
}
