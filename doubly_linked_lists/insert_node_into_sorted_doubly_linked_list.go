package doubly_linked_lists

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

func sortedInsert(listHead *DoublyLinkedListNode, dataToInsert int32) *DoublyLinkedListNode {
	newNode := &DoublyLinkedListNode{
		data: dataToInsert,
		next: nil,
		prev: nil,
	}

	if listHead == nil {
		return newNode
	}

	// handle if newNode.data is smaller than listHead.data, then newNode is new head,
	// and it will point to old head, and old head will point backwards to the new node.
	if dataToInsert <= listHead.data {
		newNode.next = listHead
		listHead.prev = newNode

		return newNode
	}

	// in case none of above: find the place where newNode should be inserted
	currentNode := listHead

	// find the node that should come before newNode (or go to end of list if none are bigger)
	for currentNode.next != nil && currentNode.next.data < dataToInsert {
		currentNode = currentNode.next
	}

	// insertion time: now currentNode is the node that should be before newNode,
	// find the node that should come after & note it to not lose rest of list.
	nodeAfterNewNode := currentNode.next

	newNode.next = nodeAfterNewNode
	newNode.prev = currentNode
	currentNode.next = newNode

	// point the next node backwards to the new node
	if nodeAfterNewNode != nil {
		nodeAfterNewNode.prev = newNode
	}

	return listHead
}
