package linked_lists

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	var previousNode *SinglyLinkedListNode = nil
	var currentNode *SinglyLinkedListNode = llist
	var nextNode *SinglyLinkedListNode = nil

	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	return previousNode
}
