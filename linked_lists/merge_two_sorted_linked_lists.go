package linked_lists

func mergeLists(headA *SinglyLinkedListNode, headB *SinglyLinkedListNode) *SinglyLinkedListNode {
	// check if either is nil, then it will be easy.
	if headA == nil {
		return headB
	}
	if headB == nil {
		return headA
	}

	// which is new head? the smaller one. assign to mergedListHead
	// in this check i move the respective lists' head to 'next' because else i will duplicate that data later.
	// the other, bigger head, no need to go to next yet.
	var mergedListHead *SinglyLinkedListNode
	if headA.data <= headB.data {
		mergedListHead = headA
		headA = headA.next
	} else {
		mergedListHead = headB
		headB = headB.next
	}

	// merged list tail will always be the end of the new list, which in fact at this point will be the head.
	mergedListTail := mergedListHead

	// now time to keep comparing and finding the smaller of the two nodes.
	// attach the smaller one to the tail, and move pointer forward if it has been attached
	for headA != nil && headB != nil {
		if headA.data <= headB.data {
			mergedListTail.next = headA
			headA = headA.next
		} else {
			mergedListTail.next = headB
			headB = headB.next
		}
		// advance to the node we just added, this is the new tail now
		mergedListTail = mergedListTail.next
	}

	// if the lists were not the same length, then attach the longer one just the rest of the list, since is already sorted
	if headA != nil {
		mergedListTail.next = headA
	} else if headB != nil {
		mergedListTail.next = headB
	}

	return mergedListHead
}
