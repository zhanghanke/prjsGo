/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    /*if l2 == nil || l2.Val == 0 {
		return l1
	}

	if l1 == nil || l1.Val == 0 {
		return l2
	}*/
	np1 := l1
	list1 := list.New()
	for np1 != nil {
		list1.PushBack(np1.Val)
		np1 = np1.Next

	}
	np2 := l2
	list2 := list.New()
	for np2 != nil {
		list2.PushBack(np2.Val)
		np2 = np2.Next

	}
	sum := 0
	idx := 0
	for e := list1.Back(); e != nil; e = e.Prev() {
		sum += int(math.Pow10(idx)) * e.Value.(int)
		idx++
	}
	idx = 0
	for e := list2.Back(); e != nil; e = e.Prev() {
		sum += int(math.Pow10(idx)) * e.Value.(int)
		idx++
	}
	nd := ListNode{}
	nd.Val = sum % 10
	nd.Next = nil

	ndptr := &nd
	sum /= 10
	for sum != 0 {
		n := ListNode{}
		n.Val = sum % 10
		//fmt.Println(n.Val)
		ndptr.Next = &n
		ndptr = &n
		sum /= 10
	}
	return &nd   
}
