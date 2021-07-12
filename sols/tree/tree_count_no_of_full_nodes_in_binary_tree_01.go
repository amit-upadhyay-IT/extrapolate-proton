package tree

import "strconv"

/* NOTE BEGIN
 * Full node in a binary tree is the one which has both the children present, i.e. left and right child both are present
 * Again, to solve this we can simply traverse the tree and maintain a count of nodes which has both their child present
 * But, we should always try to write a method to solve exactly that problem in a way that programming recursion stack
 * would return the solution value for the sub-problems and using which we will construct the final answer.
 * So, we would try to define our recursive equation in that way only.
 * FullNode(root) = 1 + FullNode(root.left) + FullNode(root.right); if root.left != nil and root.right != nil
                  = FullNode(root.left) + FullNode(root.right); if root.left == nil || root.right == nil
                  = 0 ; root.left == nil and root.right == nil
NOTE END */


func Tree_count_no_of_full_nodes_in_binary_tree_01() string {
	tree, _ := GetDummyBinaryTree()
	return strconv.Itoa(countFullNodes(tree))
}

func countFullNodes(node *TreeNode) int {
	// base case
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 0
	}
	leftFullNodes := countFullNodes(node.Left)
	rightFullNodes := countFullNodes(node.Right)
	if node.Left == nil || node.Right == nil {
		return leftFullNodes + rightFullNodes
	} else {
		return 1 + leftFullNodes + rightFullNodes
	}
}
