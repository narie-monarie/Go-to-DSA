/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    if node == nil{
        return
    }
    
    temp := node
    for temp.Next.Next != nil{
        temp.Val = temp.Next.Val
        temp = temp.Next
    }
    
    temp.Val = temp.Next.Val
    temp.Next = nil
}