/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, Val int) *TreeNode {
    if(root == nil){
        return nil
    }
    
    if(root.Val == Val){
        return root;
    }
    
    if(root.Val > Val){
        return searchBST(root.Left, Val)
    }
    
    return searchBST(root.Right, Val)
    
}