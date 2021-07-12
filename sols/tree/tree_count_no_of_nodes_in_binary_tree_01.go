package tree

import "strconv"

/* NOTE BEGIN
 * How do you count the nodes in a binary tree?
 * I will first count count the nodes in left child of root and then count right child of root and then finally
 * sum these two values and add one(counting root node too).
 * As, you can see we can recursively count the number of nodes, The recursive equation would be:
 * nodesCount(root) = 1 + nodeCount(root.left) + nodeCount(root.right)
                    = 0 ; if root is nil
NOTE END */


func Tree_count_no_of_nodes_in_binary_tree_01() string {
	tree, _ := GetDummyBinaryTree()
	return strconv.Itoa(countNodes(tree))
}

func countNodes(node *TreeNode) int {
	// base case
	if node == nil {
		return 0
	}
	leftNodes := countNodes(node.Left)
	rightNodes := countNodes(node.Right)
	return 1 + leftNodes + rightNodes
}
