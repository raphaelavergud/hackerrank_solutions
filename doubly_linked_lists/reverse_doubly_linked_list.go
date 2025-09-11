package doubly_linked_lists

func reverseDoublyLinkedList(listHead *DoublyLinkedListNode) *DoublyLinkedListNode {
	if listHead == nil || listHead.next == nil {
		return listHead
	}

	currentNode := listHead
	var newHead *DoublyLinkedListNode = nil

	for currentNode != nil {
		nextNodeInOriginalList := currentNode.next

		// the swap
		currentNode.prev, currentNode.next = currentNode.next, currentNode.prev

		newHead = currentNode                // update newHead after each swap before traversing, in case it's the last loop
		currentNode = nextNodeInOriginalList // traverse
	}

	return newHead
}
