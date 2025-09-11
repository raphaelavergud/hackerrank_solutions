package linked_lists

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func insertNodeAtHead(listHead *SinglyLinkedListNode, dataToInsert int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{
		data: dataToInsert,
		next: listHead,
	}
	return newNode // return new head
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	countStr, _ := reader.ReadString('\n')
	countStr = strings.TrimSpace(countStr)
	count, _ := strconv.Atoi(countStr)

	var list *SinglyLinkedListNode = nil // points to the head of the linked list always, its nil at first because the list is empty at first

	i := 0
	for i < count {
		dataStr, _ := reader.ReadString('\n')
		dataStr = strings.TrimSpace(dataStr)
		data64, _ := strconv.ParseInt(dataStr, 10, 32)
		data := int32(data64)

		list = insertNodeAtHead(list, data) // point to new head which will then head to the entire linked list

		i = i + 1
	}

	printLinkedList(list)
}
