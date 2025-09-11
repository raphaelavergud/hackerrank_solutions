package linked_lists

func compareLists(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) int32 {
	currentNode1 := head1
	currentNode2 := head2

	for currentNode1 != nil && currentNode2 != nil {
		if currentNode1.data != currentNode2.data {
			return 0
		}
		currentNode1 = currentNode1.next
		currentNode2 = currentNode2.next
	}
	if currentNode1 == nil && currentNode2 == nil {
		return 1
	} else {
		return 0
	}
}
