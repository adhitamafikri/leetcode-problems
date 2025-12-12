package main

import "fmt"

// Pseudo code
// 1. We have List 1 and List 2, each containing sorted value
// 2.

type ListNode struct {
	Val  int
	Next *ListNode
}

func traverse(head *ListNode) {
	if head != nil {
		fmt.Println("Item from List: ", head.Val)
		traverse(head.Next)
	}
}

func getNode(head *ListNode) (current *ListNode, next *ListNode) {
	if head != nil {
		current = head
		next = head.Next
		return
	}

	return
}

func appendList(head *ListNode, newValue int) *ListNode {
	if head == nil {
		head = &ListNode{
			Val:  newValue,
			Next: nil,
		}
		return head
	}

	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newList *ListNode

	iterate := true
	head1 := list1
	head2 := list2
	for iterate {
		c1, n1 := getNode(head1)
		c2, n2 := getNode(head2)

		// move the heads to the next
		head1 = n1
		head2 = n2

		// append our c's to the new list
		appendList(newList, c1.Val)
		appendList(newList, c2.Val)

		if n1 == nil && n2 == nil {
			iterate = false
		}
	}

	return newList
}

func traverseTwoList(list1 *ListNode, list2 *ListNode) {
	traverse(list1)
	traverse(list2)

	n1v1, n1v2 := getNode(list1)
	n2v1, n2v2 := getNode(list2)

	fmt.Println("n1v1, n1v2", n1v1, n1v2)
	fmt.Println("n2v1, n2v2", n2v1, n2v2)
}

func case1() {
	fmt.Println("Case #1\n----------")
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	// result := mergeTwoLists(&list1, &list2)
	// fmt.Println("Result: ", result)
	traverseTwoList(&list1, &list2)
	result := mergeTwoLists(&list1, &list2)
	traverse(result)
}

func case2() {
	fmt.Println("Case #2\n----------")
	list1 := ListNode{
		Val:  -101,
		Next: nil,
	}
	list2 := ListNode{
		Val:  -101,
		Next: nil,
	}
	// result := mergeTwoLists(&list1, &list2)
	// fmt.Println("Result: ", result)
	traverseTwoList(&list1, &list2)
}

func case3() {
	fmt.Println("Case #3\n----------")
	list1 := ListNode{
		Val:  -101,
		Next: nil,
	}
	list2 := ListNode{
		Val:  0,
		Next: nil,
	}
	// result := mergeTwoLists(&list1, &list2)
	// fmt.Println("Result: ", result)
	traverseTwoList(&list1, &list2)
}

func main() {
	fmt.Println("21 Merge Two Sorted Lists\n----------")
	fmt.Println("---")
	case1()
	// fmt.Println("---")
	// case2()
	// fmt.Println("---")
	// case3()
}
