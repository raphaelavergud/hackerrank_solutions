package linked_lists

func findMergeNode(headA *SinglyLinkedListNode, headB *SinglyLinkedListNode) int32 {
	// find out the length of both lists
	var lengthA int32
	currentNodeA := headA
	for lengthA = 0; currentNodeA != nil; currentNodeA = currentNodeA.next {
		lengthA = lengthA + 1
	}

	var lengthB int32
	currentNodeB := headB
	for lengthB = 0; currentNodeB != nil; currentNodeB = currentNodeB.next {
		lengthB = lengthB + 1
	}

	// back to start of each list.
	currentNodeA = headA
	currentNodeB = headB

	// check which list is longer, and from that list, subtract the length of the other list.
	// move the respective pointer of the longer list that many positions ahead
	// then, both pointers will be same distance from end of their lists
	if lengthA > lengthB {
		difference := lengthA - lengthB
		var i int32
		for i = 0; i < difference; i++ {
			currentNodeA = currentNodeA.next
		}
	} else if lengthB > lengthA {
		difference := lengthB - lengthA
		var i int32
		for i = 0; i < difference; i++ {
			currentNodeB = currentNodeB.next
		}
	}

	// this loop repeats until merge point is found.
	// the merge point will be reached at the same time because we sent one pointer ahead, so we are the same distance from the end, and as such
	// we will reach the merge point together (when both pointers are the same)
	for currentNodeA != currentNodeB {
		currentNodeA = currentNodeA.next
		currentNodeB = currentNodeB.next
	}

	return currentNodeA.data
}
