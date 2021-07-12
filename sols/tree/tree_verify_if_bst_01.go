package tree

import (
	"math"
	"strconv"
)

/* NOTE BEGIN
 * How will you confirm if a binary tree is BST or not?
 * Basically, you know the property: left child node value < parent node value < right child node value
 * If you try to perform the InOrder traversal, where left child node gets printed first, then parent node and then
 * right child node, so we should always get a sorted list.
 * One more method of verifying this same could be, going to each node and verifying if that node is following BST
 * property.

 *                          8
 *                        /   \
 *                       3    10
 *                     /  \     \
 *                    1   6     14
 *                       / \    /
 *                      4  7   13

 * The InOrder traversal will result: 1,3,4,6,7,8,10,13,14 which is in sorted order, so we can say that the above tree
 * is a binary search tree

 * Another way of solving this could be defining the solution in terms of itself
 * IsBST(root) = root.val > max(root.left) and root.val < min(root.right) && IsBST(root.left) && IsBST(root.right); if
               = true; if node is leaf node
 * here IsBST(root) is a function of root of tree which will return true of false if tree is BST or not respectively

 * What would be the time complexity?
 * If you see at each node, you try to compare value and at the same time you try to find maximum or minimum in their
 * child tree, which in worst case can take O(n) time in case of skewed tree.
 * So, overall time complexity would be O(n^2)
 * We can improve this complexity, how?
 * We know that we can traverse from root to leave, now in that process, if we can pass the limit in the recursive call
 * and have a logic inside our method which will will check the limit bound of the nodes value.
 * Let's try to understand the process with an example

 *                          8
 *                        /   \
 *                       3    10
 *                     /  \     \
 *                    1   6     14
 *                       / \    /
 *                      4  7   13

 * when we are at 8, we know that maximum limit of elements in the left subtree should be 8, and the minimum limit of
 * elements in the right sub tree should also be 8.
 * Lets say now we reach the left child node of 8, i.e. we reach at 3, the maximum limit will change to 3 for left
 * child, but maximum limit will still remain same, i.e. 8 for the right child subtree.
 * So, basically, while moving to left child node, we change the maximum limit to the node value and while moving to
 * right child node, we change the minimum limit to the parent node value.
 * This will help us solve this problem in O(n) time only

NOTE END */

func Tree_verify_if_bst_01() string {
	nonBst, _ := GetDummyBinaryTree()
	res := ""
	res += strconv.FormatBool(IsBST(nonBst)) + ", "
	bst := GetDummyBST()
	res += strconv.FormatBool(IsBST(bst))
	res += "\n"
	res += strconv.FormatBool(IsBST2(nonBst, -math.MaxInt64, math.MaxInt64)) + ", "
	res += strconv.FormatBool(IsBST2(bst, -math.MaxInt64, math.MaxInt64))
	return res
}

func IsBST(node *TreeNode) bool {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return true
	}
	return node.Val > getMaxAssumingBST(node.Left) &&
		node.Val < getMinAssumingBST(node.Right) &&
		IsBST(node.Left) && IsBST(node.Right)
}


// this method finds maximum by assuming that it is BST
func getMaxAssumingBST(node *TreeNode) int {
	if node == nil {
		return -math.MaxInt64
	}
	temp := node
	for temp.Right != nil {
		temp = temp.Right
	}
	return temp.Val
}

func getMinAssumingBST(node *TreeNode) int {
	if node == nil {
		return math.MaxInt64
	}
	temp := node
	for temp.Left != nil {
		temp = temp.Left
	}
	return temp.Val
}


func IsBST2(node *TreeNode, minLimit, maxLimit int) bool {
	// base case, node value breaches the limits
	if node == nil {
		return true
	}
	if node.Val < minLimit || node.Val > maxLimit {
		return false
	}
	return IsBST2(node.Left, minLimit, node.Val-1) && IsBST2(node.Right, node.Val+1, maxLimit)
}
