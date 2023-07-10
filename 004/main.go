package main

import (
	"strconv"
)

/*
876. Middle of the Linked List
Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

Constraints:
The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100

Example 1:
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	println("Start main")
	var task string = "1-2-3-4-5"
	println(middleNode(buildList(task)).Val)
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next
	}

	return head
}

func buildList(values string) *ListNode {
	var tmpVal string
	var currentNode *ListNode
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] != '-' {
			tmpVal = string(values[i]) + tmpVal
		} else {
			//println(tmpVal)
			currentNode = createNewNode(tmpVal, currentNode)
			tmpVal = ""
		}
	}
	//println(tmpVal)
	currentNode = createNewNode(tmpVal, currentNode)
	return currentNode
}

func printList(head *ListNode) string {
	currentNode := head
	var result string
	for currentNode.Next != nil {
		result += strconv.Itoa(currentNode.Val) + "-"
		currentNode = currentNode.Next
	}
	result += strconv.Itoa(currentNode.Val)
	return result
}

func createNewNode(value string, nextNode *ListNode) *ListNode {
	var p ListNode
	p.Val, _ = strconv.Atoi(value)
	p.Next = nextNode
	return &p
}
