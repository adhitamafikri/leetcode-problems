package main

import "fmt"

type LinkedList struct {
	head  int
	tail  int
	nodes *LinkedListNode
}

type LinkedListNode struct {
	id   int
	data string
	next *LinkedListNode
}

// Recursively traverse the '.next' property of each node, until the supplied 'node' is nil.
// Which means the current iteration has no actual node value
func (l *LinkedList) traverseAndPrint(node *LinkedListNode) {
	if node != nil {
		fmt.Println("List item: ", node)
		l.traverseAndPrint(node.next)
	}
}

func (l *LinkedList) traverseAndAppend(node *LinkedListNode, newNode *LinkedListNode) {
	if node.next != nil {
		l.traverseAndAppend(node.next, newNode)
	} else {
		node.next = newNode
	}
}

func (l *LinkedList) appendNode(data string) {
	var newNode LinkedListNode
	if l.head < 0 && l.tail < 0 {
		l.head = 0
		l.tail = 0
		newNode = LinkedListNode{
			id:   0,
			data: data,
			next: nil,
		}
		l.nodes = &newNode
	} else {
		l.tail++
		newNode = LinkedListNode{
			id:   l.tail,
			data: data,
			next: nil,
		}
		l.traverseAndAppend(l.nodes, &newNode)
	}
}

func main() {
	fmt.Println("This is the linked list again")
	list := LinkedList{
		head:  -1,
		tail:  -1,
		nodes: nil,
	}
	fmt.Println("My Linked List: ", list, "\n")

	fmt.Println("\nTraversing and Appending New Item to Linked List\n---------------")
	list.appendNode("Item #1")
	fmt.Println("My Linked List: ", list, "\n")

	list.appendNode("Item #2")
	fmt.Println("My Linked List: ", list, "\n")

	list.appendNode("Item #3")
	fmt.Println("My Linked List: ", list, "\n")

	list.appendNode("Item #4")
	fmt.Println("My Linked List: ", list, "\n")

	list.appendNode("Item #5")
	fmt.Println("My Linked List: ", list, "\n")

	list.appendNode("Item #6")
	fmt.Println("My Linked List: ", list, "\n")

	list.appendNode("Item #7")
	fmt.Println("My Linked List: ", list, "\n")

	fmt.Println("\nTraversing Linked List item\n---------------")
	list.traverseAndPrint(list.nodes)
}
