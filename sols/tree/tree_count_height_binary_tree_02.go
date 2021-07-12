package tree

import "strconv"

/* NOTE BEGIN
 * Another approach towards solution could be defining a proper recursive solution like:
 * height(root) = 1 + max(height(root.left), height(root.right))
 *              = 0; if root is nil

 * You can see the previous solution and this solution and decide yourself which one is better in terms of solutioning,
 * implementation and cleanliness.
NOTE END */

func Tree_count_height_binary_tree_02() string {

	tree, _ := GetDummyBinaryTree()
	return strconv.Itoa(heightTree(tree))
}

func heightTree(node *TreeNode) int {
	if node == nil {return 0}
	if node.Left == nil && node.Right == nil {
		return 0  // height of leaf node is zero
	}
	leftHeight := heightTree(node.Left)
	rightHeight := heightTree(node.Right)
	return 1 + getMax_tree_count_height(leftHeight, rightHeight)
}

func getMax_tree_count_height(a, b int) int {
	if a > b {
		return a
	}
	return b
}
