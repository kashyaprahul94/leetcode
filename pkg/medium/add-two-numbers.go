package medium

import "fmt"

// ListNode represents a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) String() string {
	result := "Start"
	head := list

	for head != nil {
		result = fmt.Sprintf("%s -> %v", result, head.Val)
		head = head.Next
	}

	return fmt.Sprintf("%s -> End", result)
}

func (list *ListNode) toArray() []int {
	var result []int

	head := list

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val: 0,
	}
	current := result

	carry := 0

	for l1 != nil || l2 != nil {

		x, y, sum := 0, 0, 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum = x + y + carry

		if sum >= 10 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}

		current.Next = &ListNode{
			Val: sum,
		}
		current = current.Next
	}

	if carry != 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}

	return result.Next
}
