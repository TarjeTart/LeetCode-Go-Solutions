func detectCycle(head *ListNode) *ListNode {
    for head != nil {
        if head.Val == math.MaxInt {
            return head
        }

        head.Val = math.MaxInt
        head = head.Next
    }

    return nil
}
