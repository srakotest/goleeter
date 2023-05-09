package main

import "strconv"

/*
234. Palindrome Linked List
Given the head of a singly linked list, return true if it is a
palindrome or false otherwise.

Input: head = [1,2,2,1]
Output: true

Input: head = [1,2]
Output: false

Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var task string = "1-2-3"
	//printList(buildList(task))
	println(task, isPalindrome(buildList(task)))
	task = "1"
	println(task, isPalindrome(buildList(task)))

	task = "1-2"
	println(task, isPalindrome(buildList(task)))

	task = "1-1"
	println(task, isPalindrome(buildList(task)))

	task = "1-2-2-1"
	println(task, isPalindrome(buildList(task)))

	task = "1-2-3-2-1"
	println(task, isPalindrome(buildList(task)))

	task = "1-2-1"
	println(task, isPalindrome(buildList(task)))

	task = "1-2-3"
	println(task, isPalindrome(buildList(task)))

	task = "1-2"
	println(task, isPalindrome(buildList(task)))

	task = "1-2-3-4"
	println(task, isPalindrome(buildList(task)))

	task = "1-1-2-1"
	println(task, isPalindrome(buildList(task)))
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

func printList(head *ListNode) {
	currentNode := head
	println("printList")
	if currentNode == nil {
		println("List is empty")
		return
	}
	for currentNode.Next != nil {
		print(currentNode.Val, "-")
		currentNode = currentNode.Next
	}
	println(currentNode.Val)
}

func createNewNode(value string, nextNode *ListNode) *ListNode {
	var p ListNode
	p.Val, _ = strconv.Atoi(value)
	p.Next = nextNode
	return &p
}

func isPalindrome(head *ListNode) bool {
	currentNode := head
	var listLen int = 1
	//count list length
	for currentNode.Next != nil {
		listLen++
		currentNode = currentNode.Next
	}
	println("List Length is ", listLen)

	var secondHead *ListNode
	if listLen > 2 {
		//revert second part from middle
		var prevNode *ListNode
		var nextNode *ListNode
		currentNode = head
		println("revert")
		for i := 0; currentNode.Next != nil; i++ {
			//println(currentNode.Val)
			if i == listLen/2-1 {
				nextNode = currentNode.Next.Next
				currentNode.Next.Next = nil
				prevNode = currentNode.Next
				currentNode.Next = nil
				currentNode = nextNode
			} else if i > listLen/2-1 {
				nextNode = currentNode.Next
				currentNode.Next = prevNode
				prevNode = currentNode
				currentNode = nextNode
			}
			//prevNode = currentNode
			if i < listLen/2-1 {
				currentNode = currentNode.Next
			}
		}
		secondHead = currentNode
		secondHead.Next = prevNode
	} else if listLen == 2 {
		//only 2 nodes
		secondHead = head.Next
		head.Next = nil
	} else if listLen == 1 {
		return true
	}

	printList(head)
	printList(secondHead)

	// compare 2 parts one by one
	if listLen/2 >= 2 {
		currentNode = head
		secondCurrentNode := secondHead
		for currentNode.Next != nil {
			//println(currentNode.Val, secondCurrentNode.Val)
			if currentNode.Val == secondCurrentNode.Val {
				secondCurrentNode = secondCurrentNode.Next
				currentNode = currentNode.Next
			} else {
				return false
			}
		}
		if currentNode.Val != secondCurrentNode.Val {
			return false
		}
	} else if secondHead.Val == head.Val {
		return true
	} else {
		return false
	}

	return true
}

///submited version
func xisPalindrome(head *ListNode) bool {
	currentNode := head
	var listLen int = 1
	//count list length
	for currentNode.Next != nil {
		listLen++
		currentNode = currentNode.Next
	}

	var secondHead *ListNode
	if listLen > 2 {
		//revert second part from middle
		var prevNode *ListNode
		var nextNode *ListNode
		currentNode = head

		for i := 0; currentNode.Next != nil; i++ {

			if i == listLen/2-1 {
				nextNode = currentNode.Next.Next
				currentNode.Next.Next = nil
				prevNode = currentNode.Next
				currentNode.Next = nil
				currentNode = nextNode
			} else if i > listLen/2-1 {
				nextNode = currentNode.Next
				currentNode.Next = prevNode
				prevNode = currentNode
				currentNode = nextNode
			}
			//prevNode = currentNode
			if i < listLen/2-1 {
				currentNode = currentNode.Next
			}
		}
		secondHead = currentNode
		secondHead.Next = prevNode
	} else if listLen == 2 {
		//only 2 nodes
		secondHead = head.Next
		head.Next = nil
	} else if listLen == 1 {
		return true
	}

	// compare 2 parts one by one
	if listLen/2 >= 2 {
		currentNode = head
		secondCurrentNode := secondHead
		for currentNode.Next != nil {
			if currentNode.Val == secondCurrentNode.Val {
				secondCurrentNode = secondCurrentNode.Next
				currentNode = currentNode.Next
			} else {
				return false
			}
		}

		if currentNode.Val != secondCurrentNode.Val {
			return false
		}
	} else if secondHead.Val == head.Val {
		return true
	} else {
		return false
	}

	return true
}

func isPalindromeGpt(head *ListNode) bool {
	// Сперва определим длину списка.
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}

	// Затем разделим список на две части. Если длина списка нечетная, то
	// средний элемент будет находиться в первой части.

	node = head
	for i := 0; i < length/2; i++ {
		node = node.Next
	}

	// Развернем вторую часть списка.
	var prev *ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	// Сравним первую и вторую части списка.
	for i := 0; i < length/2; i++ {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}

	// Если мы дошли до этой точки, то список является палиндромом.
	return true
}
