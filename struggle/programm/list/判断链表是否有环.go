package list

//leetCode 141
//链表中节点的数目范围是 [0, 104]
//-105 <= ListNode.val <= 105
//pos 为 -1 或者链表中的一个 有效索引 。

//1.暴力解法
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	i := 0
	for head.Next != nil {
		head = head.Next
		i++
		if i > 10000 {
			return true
		}

	}
	return false
}

//2.hash法
func hasCycle2(head *ListNode) bool {
	//seen := make(map[*ListNode]int)
	seen := map[*ListNode]struct{}{}
	if head == nil {
		return false
	}
	for head.Next != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

//3.快慢指针
func hasCycle3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
