// A linked list is a linear data structure, in which the elements are not stored at contiguous memory locations.
// The elements in a linked list are linked using pointers.
// A linked list consists of nodes where each node contains a data field and a reference(link) to the next node in the list.

package main

import "fmt"

// Node represents a single element in the linked list.
type Node struct {
	data int
	next Node
}

// LinkedList represents the linked list data structure.
type LinkedList struct {
	head Node
}

// Append adds a new node with the given data to the end of the linked list.
func (ll *LinkedList) Append(data int) {
	newNode := Node{data: data}

	if ll.head.data == 0 {
		ll.head = newNode
		return
	}

	currentNode := ll.head
	for currentNode.next.data != 0 {
		currentNode = currentNode.next
	}

	currentNode.next = newNode
}

// PrintList prints the linked list elements.
func (ll *LinkedList) PrintList() {
	currentNode := ll.head
	for currentNode.data != 0 {
		fmt.Printf("%d -> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}

func main() {
	ll := LinkedList{}

	// Append elements to the linked list
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	// Print the linked list
	ll.PrintList()
}
