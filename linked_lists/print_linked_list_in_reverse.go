package linked_lists

import "fmt"

func reversePrint(llist *SinglyLinkedListNode) {
	var values []int32

	for currentNode := llist; currentNode != nil; currentNode = currentNode.next {
		values = append(values, currentNode.data)
	}

	reverseIndex := len(values) - 1
	for reverseIndex >= 0 {
		valueToPrint := values[reverseIndex]
		fmt.Println(valueToPrint)
		reverseIndex = reverseIndex - 1
	}
}
