package tree

import "strconv"

/* NOTE BEGIN
 * In the same fashion of counting leaf nodes in a binary tree, we can count the non-leaf nodes in the binary tree.
 * Can we define the solution recursively, i.e. is the solution self-similar?
 * Yes, we can say that total non leaf nodes is non of non-leaf nodes in left child plus number of non-leaf nodes in
 * right child plus one.

 * NonLeafCount(root) = NonLeafCount(root.left) + NonLeafCount(root.right) + 1
                      = 0 ; if root.left is nil and root.right is nil
NOTE END */


func Tree_count_no_of_nonleaf_nodes_in_binary_tree_01() string {
	root, _ := GetDummyBinaryTree()
	return strconv.Itoa(nonLeafCount(root))
}

func nonLeafCount(node *TreeNode) int {
	// base condition
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 0
	}
	leftNonLeaf := nonLeafCount(node.Left)
	rightNonLeaf := nonLeafCount(node.Right)
	return 1 + leftNonLeaf + rightNonLeaf
}
