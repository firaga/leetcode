package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := getLinkedList()
	newHead := reverseList(head)
	fmt.Println(newHead)
}

func getLinkedList() *ListNode {

	var head *ListNode
	data := [5]int{1, 2, 3, 4, 5}
	var prev *ListNode
	for k, v := range data {
		tmp := &ListNode{
			Val:  v,
			Next: nil,
		}
		if k == 0 {
			head = tmp
		}
		if prev != nil {
			prev.Next = tmp
		}
		prev = tmp
	}
	return head
}
