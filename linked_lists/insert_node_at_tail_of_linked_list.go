package linked_lists

func insertNodeAtTail(head *SinglyLinkedListNode, data int32) {
	newNode := &SinglyLinkedListNode{
		data: data,
		next: nil,
	}

	if head == nil {
		head = newNode
	} else {
		currentNode := head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}
