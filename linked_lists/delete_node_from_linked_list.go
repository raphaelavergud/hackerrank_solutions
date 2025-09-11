package linked_lists

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	if position == 0 {
		newHead := llist.next
		return newHead
	}

	nodeBeforePosition := llist

	for currentPosition := 0; currentPosition < int(position)-1; currentPosition++ {
		nodeBeforePosition = nodeBeforePosition.next
	}

	nodeToDelete := nodeBeforePosition.next

	nodeBeforePosition.next = nodeToDelete.next

	return llist
}
