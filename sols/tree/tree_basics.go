package tree

/* NOTE BEGIN
 * Binary Tree: any tree which has got at-most two children is called binary tree.
 * Tree Traversal:
        * In-order
        * Pre-order
        * Post-order
        * Level-order
        * DFS and BFS can be categorized in above categories
 * These traversal algorithms were initially proposed for traversing binary tree, but some extension to them can allow
 * the traversal on n-array tree.

 * Example of binary tree:
 *       root
 *      /    \
 *   left    right

 * In Order Traversal:
     * Left, Root, Right : note that root is in middle

 * Pre Order Traversal:
     * Root, Left, Right : note that root is first

 * Post Order Traversal:
     * Left, Right, Root : note that root is at end

 * In all the above traversals, left child node is always visited before right child node

------------------------------------------------------------------------------------------------------------------------
TIP: quick way to read preorder, inorder and postorder given a tree structure
------------------------------------------------------------------------------------------------------------------------
 *      A
 *    /  \
 *   B   C

 * Assume that every node in the tree have a children, even if it is left node, i.e. make dummy leaf node
 * Example:
 *      A
 *    /  \
 *   B   C
 *  / \ / \

 * If we start walking the tree from root, we will visit every node 3 times.
 * PreOrder: Print the node value whenever you visit first time, i.e. A,B,C is preorder of above tree
 * InOrder: Print the node value whenever you visit second time, i.e. B,A,C is in order
 * PostOrder: Print the node value whenever you visit the node third time, i.e. B,C,A is postorder of above tree

 * Example:
 *                    A
 *                  /   \
 *                 B     C
 *               / \    / \
 *                 D   E
 *                / \ / \
 *               F       G
 *              / \     / \

 * PreOrder : A,B,D,F,C,E,G
 * InOrder  : B,F,D,A,E,G,C
 * PostOrder: F,D,B,G,E,C,A

------------------------------------------------------------------------------------------------------------------------
Recursive implementation of traversal
------------------------------------------------------------------------------------------------------------------------
1) InOrder:
 * We know that in this case, the root value gets printed after the left child node value, so we will first explore all
 * the left child nodes, after we are done exploring we will print their values, i.e. first left leaf node in the tree
 * will get printed, now that the exploration of left child nodes are done, program will backtrack it self. In the
 * process of backtracking we will keep on printing the node values.
 * Now, after backtracking is done, we want to explore the right nodes as well and the exploration needs to be done in
 * the same fashion as we have done for the root.

 * program would look something like this:
func InOrder(node *Tree) {
    if node == nil {
        return
    }
    InOrder(node.Left)
    print(node.Value)
    InOrder(node.Right)
}

* Time complexity: O(n) as each node is explored once and each visit of node we perform constant time operation
* Space complexity: O(n) in worst case we can have skewed tree and in that case n function stack will be created during
* program execution.

------------------------------------------------------------------------------------------------------------------------
Number of binary trees possible
------------------------------------------------------------------------------------------------------------------------
 * Number of unlabeled tree possible with n nodes, un-labeled tree refers to just tree structure without having values
 * inside tree nodes

 * With one node, only one tree is possible
 * With two nodes, two types of tree can be formed
 *        N                N
 *       /       and        \
 *      N                   N

 * With three nodes, 5 kind of trees can be formed
 *        N              N        N              N            N
 *       / \            /         \             /              \
 *      N  N           N           N           N                N
 *                    /            \            \              /
 *                   N             N            N             N

 * There is a problem discussed in this section below where I have written a recursive solution to find it.
 * However, you can use a formula to calculate it (2nCn)/(n+1)

 * In can of possible trees if nodes are labelled, we know that each structure will be able to generate n! trees.

NOTE END */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
                     1
                   /    \
                  2      3
                /  \    / \
               4   5   6  7
              / \  /\
             8  9 10 11
returns the above tree and count of nodes in tree
*/
func GetDummyBinaryTree() (*TreeNode, int) {
	eleven := &TreeNode{Val: 11, Left: nil, Right: nil}
	ten := &TreeNode{Val: 10, Left: nil, Right: nil}
	nine := &TreeNode{Val: 9, Left: nil, Right: nil}
	eight := &TreeNode{Val: 8, Left: nil, Right: nil}
	seven := &TreeNode{Val: 7, Left: nil, Right: nil}
	six := &TreeNode{Val: 6, Left: nil, Right: nil}
	five := &TreeNode{Val: 5, Left: ten, Right: eleven}
	four := &TreeNode{Val: 4, Left: eight, Right: nine}
	three := &TreeNode{Val: 3, Left: six, Right: seven}
	two := &TreeNode{Val: 2, Left: four, Right: five}
	root := &TreeNode{Val: 1, Left: two, Right: three}
	return root, 11
}

/*
        1
      /  \
     2    3
    / \
   4   5
        \
         6
*/
func GetDummyBinaryTree2() *TreeNode {

	six := &TreeNode{Val: 6, Left: nil, Right: nil}
	five := &TreeNode{Val: 5, Left: nil, Right: six}
	four := &TreeNode{Val: 4, Left: nil, Right: nil}
	three := &TreeNode{Val: 3, Left: nil, Right: nil}
	two := &TreeNode{Val: 2, Left: four, Right: five}
	root := &TreeNode{Val: 1, Left: two, Right: three}
	return root
}

func GetDummyBST() *TreeNode {
	one := &TreeNode{Val: 1, Left: nil, Right: nil}
	four := &TreeNode{Val: 4, Left: nil, Right: nil}
	seven := &TreeNode{Val: 7, Left: nil, Right: nil}
	thirteen := &TreeNode{Val: 13, Left: nil, Right: nil}
	six := &TreeNode{Val: 6, Left: four, Right: seven}
	fourteen := &TreeNode{Val: 14, Left: thirteen, Right: nil}
	three := &TreeNode{Val: 3, Left: one, Right: six}
	ten := &TreeNode{Val: 10, Left: nil, Right: fourteen}
	root := &TreeNode{Val: 8, Left: three, Right: ten}
	return root
}
