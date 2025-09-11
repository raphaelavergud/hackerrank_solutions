package linked_lists

func removeDuplicates(listHead *SinglyLinkedListNode) *SinglyLinkedListNode {
	currentNode := listHead
	for currentNode != nil {
		checkNextNodeHelper := currentNode.next
		// loop ends when either helpernode is empty, or next node is different from current node
		// until then, traverse
		for checkNextNodeHelper != nil && currentNode.data == checkNextNodeHelper.data {
			// in here only takes place when currentnode and nextnode are duplicates.
			checkNextNodeHelper = checkNextNodeHelper.next
		}

		// in case previous loop ever took place, we will skip ahead to wherever the nextnodehelper went
		currentNode.next = checkNextNodeHelper

		// traverse to the next node and compare again. until this loop ends.
		currentNode = checkNextNodeHelper
	}
	return listHead
}
