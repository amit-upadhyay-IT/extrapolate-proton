package tree

import (
	"strconv"
)

/* NOTE BEGIN
 * How will you count the height of a binary tree intuitively?
 * Basically, height of a tree is the largest path from the root node to any leaf node.

 * This time we will try to solve this problem in two different ways, both would be recursive only but then after you
 * understand both the solution, you will be able to figure out why defining a solution in itself is important. NOTE, by
 * defining a solution in itself I mean that each of the sum-problem solution should return some answer, using which
 * the construction of root solution will be given.

 * One, native approach could be traversing over the tree and in this process maintain a set of shortest number of edges
 * visited whenever you reach any leaf node. And you final answer will the maximum among this maintained set.

 * Another approach towards solution could be defining a proper recursive solution like:
 * height(root) = 1 + max(height(root.left), height(root.right))
 *              = 0; if root is nil

 * Let's first implement the first approach, which in my opinion is bad solutioning for the problem.
NOTE END */


func Tree_count_height_binary_tree_01() string {

	dumTree, _ := GetDummyBinaryTree()
	allLeafDistance := make([]int, 0)
	getHeightBad(dumTree, 0, &allLeafDistance)
	//fmt.Println(allLeafDistance)
	allLeafDistance = make([]int, 0)
	maxHeight = 0
	dumTree2 := GetDummyBinaryTree2()
	getHeightBad(dumTree2, 0, &allLeafDistance)
	//fmt.Println(allLeafDistance)
	return strconv.Itoa(maxHeight)
}

var maxHeight int

// maintains shortest edge visited count
func getHeightBad(node *TreeNode, currentHeight int, set *[]int) {

	if node == nil {
		return
	}
	// base condition: i.e. when you decide on a possible result
	if node.Left == nil && node.Right == nil {
		*set = append(*set, currentHeight)
		if currentHeight > maxHeight {
			maxHeight = currentHeight
		}
	}

	if node.Left != nil {
		// increase the height, coz we are moving from current node to its child node
		currentHeight++
		getHeightBad(node.Left, currentHeight, set)
		// during backtrack, we need to decrease the height, because current node is parent of node.left
		currentHeight--
	}

	if node.Right != nil {
		currentHeight++
		getHeightBad(node.Right, currentHeight, set)
		currentHeight--
	}
}


