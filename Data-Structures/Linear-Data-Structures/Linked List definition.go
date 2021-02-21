// Linked List Definition.
package main

type ListNode struct { // Linked List Node
	data interface{} // data
	next *ListNode   // pointer to the next ListNode
}

// A LinkedList contains pointer to first node (head) and size stores size of linked list.
type LinkedList struct {
	head *ListNode
	size int
}
