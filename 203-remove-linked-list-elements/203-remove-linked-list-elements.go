/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val{
        head = head.Next
    }
    
    temp := head
    
    for temp != nil && temp.Next != nil{
       
        if temp.Next.Val == val{
            temp.Next = temp.Next.Next
        }else{
        temp = temp.Next
        }
    }
    return head
}