package tree

import "strconv"

/* NOTE BEGIN
 * How do you count the number of leaf node?
 * To count, you need to confirm if a node is leaf node or not, so how do you confirm if a node is leaf?
 * If the left and right child of the node are empty, then we can say that its a leaf node.

 * To solve this programatically, we need to visit each node and see if that node is leaf node or not, i.e. we can do
 * tree traversal and keep a count of leaf nodes, increment the count whenever you encounter the leaf node in the
 * process of traversal.

 * Or, you can recursively define the solution in this form:
 * leafCount(root) = leafCount(root.left) + leafCount(root.right)
                   = 1; if root.left == nil and root.right == nil
 * This above recurrence relation looks more intuitive, as we can clearly define the function solving subproblems and
 * using result of those solution to form final solution.
NOTE END */

func Tree_count_no_of_leaf_nodes_in_binary_tree_01() string {
	root, _ := GetDummyBinaryTree()
	return strconv.Itoa(countLeafNodes(root))
}

func countLeafNodes(node *TreeNode) int {
	// base case
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	leftLeafNodes := countLeafNodes(node.Left)
	rightLeafNodes := countLeafNodes(node.Right)
	return leftLeafNodes + rightLeafNodes
}
