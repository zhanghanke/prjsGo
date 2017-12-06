/**
* 第一次，题目理解错误。再次感叹自己的水平实太水了..
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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

/*
整型溢出，即使是把sum定义为int64还是没办法解决，只能改其他办法
Input:
[2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9]
[5,6,4,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9,9,9,9]
Output:
[7,0,8,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4]
Expected:
[7,0,8,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,1,4,3,9,1]
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	idx := 0
	for l1 != nil {
		sum += int(math.Pow10(idx)) * l1.Val
		idx++
		l1 = l1.Next
	}
	idx = 0
	for l2 != nil {
		sum += int(math.Pow10(idx)) * l2.Val
		idx++
		l2 = l2.Next
	}
	if sum == 0 {
		return &ListNode{Val: 0, Next: nil}
	}
	var head *ListNode = &ListNode{}
	head.Val = -1
	head.Next = nil
	ndptr := head
	for sum != 0 {
		n := ListNode{}
		n.Val = sum % 10
		ndptr.Next = &n
		ndptr = &n
		sum /= 10
	}
	return head.Next
}

/**
* 去掉中间的sum累和的做法，PS:做了几道题之后，我以后不敢自称是程序员了
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = &ListNode{}
	head.Val = -1
	head.Next = nil
	ndptr := head
	needPlus := false
	for l1 != nil || l2 != nil || needPlus {
		n := ListNode{}
		ndptr.Next = &n
		ndptr = &n
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if needPlus {
			sum++
			needPlus = false
		}
		n.Val = sum % 10
		if sum >= 10 {
			needPlus = true
		}
	}
	return head.Next
