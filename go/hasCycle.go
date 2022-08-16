package main

func hasCycle(head *ListNode) bool {

 	if head == nil { return false }

	slow := head.Next 

	if slow == nil { return false }
	
	fast := slow.Next 

	for fast != nil && slow != nil {
		if fast == slow { return true }

		slow = slow.Next
		fast = fast.Next

		if fast != nil  { fast = fast.Next }
	}   

	return false
}
