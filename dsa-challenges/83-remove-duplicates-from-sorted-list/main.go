package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func traverse(node *ListNode) *ListNode {
	if node != nil {
		fmt.Println("The item: ", node)
		traverse(node.Next)
	}

	return node
}

// The Linked List is guaranteed to be SORTED
var prevHead *ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil {
		if prevHead == nil {
			prevHead = head
			deleteDuplicates(prevHead.Next)
		} else if prevHead.Val != head.Val {
			prevHead = head
			deleteDuplicates(head.Next)
		} else {
			prevHead.Next = head.Next
			deleteDuplicates(head.Next)
		}
	}

	prevHead = nil
	return head
}

func main() {
	fmt.Println("This is the challenge, tweak the example list variable below")
	case1()
	case2()
	case3()
	case4()
	case5()
	case6()
	case7()
}

func case1() {
	fmt.Println("Case #1")
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val:  6,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("Iterating Nodes in the List\n--------------")
	traverse(&list)

	fmt.Println("\n\nDeleting Duplicated Items from the List\n--------------")
	newList := deleteDuplicates(&list)

	fmt.Println("\n\nNew List After Deleting Duplicate Items\n--------------")
	traverse(newList)
}

func case2() {
	fmt.Println("Case #2")
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
	}

	fmt.Println("Iterating Nodes in the List\n--------------")
	traverse(&list)

	fmt.Println("\n\nDeleting Duplicated Items from the List\n--------------")
	newList := deleteDuplicates(&list)

	fmt.Println("\n\nNew List After Deleting Duplicate Items\n--------------")
	traverse(newList)
}

func case3() {
	fmt.Println("Case #3")
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}

	fmt.Println("Iterating Nodes in the List\n--------------")
	traverse(&list)

	fmt.Println("\n\nDeleting Duplicated Items from the List\n--------------")
	newList := deleteDuplicates(&list)

	fmt.Println("\n\nNew List After Deleting Duplicate Items\n--------------")
	traverse(newList)
}

func case4() {
	fmt.Println("Case #4")
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}

	fmt.Println("Iterating Nodes in the List\n--------------")
	traverse(&list)

	fmt.Println("\n\nDeleting Duplicated Items from the List\n--------------")
	newList := deleteDuplicates(&list)

	fmt.Println("\n\nNew List After Deleting Duplicate Items\n--------------")
	traverse(newList)
}

func case5() {
	fmt.Println("Case #5")
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	fmt.Println("Iterating Nodes in the List\n--------------")
	traverse(&list)

	fmt.Println("\n\nDeleting Duplicated Items from the List\n--------------")
	newList := deleteDuplicates(&list)

	fmt.Println("\n\nNew List After Deleting Duplicate Items\n--------------")
	traverse(newList)
}

func case6() {
	fmt.Println("Case #6")
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	fmt.Println("Iterating Nodes in the List\n--------------")
	traverse(&list)

	fmt.Println("\n\nDeleting Duplicated Items from the List\n--------------")
	newList := deleteDuplicates(&list)

	fmt.Println("\n\nNew List After Deleting Duplicate Items\n--------------")
	traverse(newList)
}

func case7() {
	fmt.Println("Case #7")
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	fmt.Println("Iterating Nodes in the List\n--------------")
	traverse(&list)

	fmt.Println("\n\nDeleting Duplicated Items from the List\n--------------")
	newList := deleteDuplicates(&list)

	fmt.Println("\n\nNew List After Deleting Duplicate Items\n--------------")
	traverse(newList)
}
