package list

//剑指offer 52
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := 0
	b := 0
	for headA != nil {
		headA = headA.Next
		a++
	}
	for headB != nil {
		headB = headB.Next
		b++
	}
	curA, curB := headA, headB
	if a > b {
		for a != b {
			curA = curA.Next
			a--
		}
	}
	if b > a {
		for a != b {
			curB = curB.Next
			b--
		}
	}
	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}
