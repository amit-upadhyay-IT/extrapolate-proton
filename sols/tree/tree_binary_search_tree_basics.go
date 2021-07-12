package tree

import (
	"fmt"
	"strings"
)

/* NOTE BEGIN
 * Binary Search Tree is a kind of binary tree with some added properties, the added properties help in searching faster
 * in the tree.
 * Property of BST:
     * Left child node has less value than the parent node value.
     * Right child node has high value than the parent node value.

 * Mostly, BSTs (or particularly balanced BSTs) are used for storing the keys in key-value pair, so we assume that in
 * BST the numbers are distinct.

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
 * There is a different question under this topic, you can read that to know more solutions for this problem.

------------------------------------------------------------------------------------------------------------------------
Construction of BST
------------------------------------------------------------------------------------------------------------------------
 * To construct, we need to assume a root and consider that as a BST (a single node in tree is always BST), and insert
 * the items in the tree.
 * How, do we insert item in tree?
 * Basically, for every insert operation we need to create a new node and put it as the left child or the right child
 * of some leaf node, right? We know that we can only put new nodes as child of leaf node, otherwise to put them in
 * middle, we will need to re-arrange the nodes. If we do re-arrangements it would be even better, but for now, lets
 * just simply insert the items in the order they are coming.

 * To insert:
     * create new node with input value
     * reach to appropriate leaf node depending on the input value
     * set the node either as left child or right child depending of input value
 * Recursive equation:
     * We need to define the insert method, it should perform one and only one task
     * should the method return something or not? this will depend on the solution we think
     * ins(root, val) => root = node(val) ; if root is nil
                      => root.left = ins(root.left, val) if val < root.left
                      => root.right = ins(root.right, val) if val > root.right

     * If you ask me why are you returning values from the insert method, I will say I want to assign the newly inserted
     * node as either the left child or the right child of the parent node. And in the single method call I do not want
     * to store the parent node and child node both, so I return the current node from the method and the recursion
     * call stack help me to assign that node to either left or right of the parent node, as once node is returned from
     * method, the top method from recursion call stack will get poped off, and now top method on stack will be the
     * parent of the previously returned node.

------------------------------------------------------------------------------------------------------------------------
Delete an item from BST
------------------------------------------------------------------------------------------------------------------------
 * The best way to write deletion algorithm is to try to delete items from a BST.
 * If you try yourself, you will be able to categorize the deletion process into 3 categories:
     * Deleting leaf node
     * Deleting node with either left child or right child
     * Deleting node having both child

 * To Delete leaf node, just reach the parent of the node and mark that respective child node as empty
 * To Delete node with just one child, get the address of child node of that node, and save that into the parent node
 * To Delete node with both child, get either inorder successor or inorder predecessor and swap the values, note that
     * after this you need to delete that node with which you have swapped the value. And deleting that node will be
     * easy, coz that would fall under either case 1 or case 2 as above described.

NOTE END */


func Tree_binary_search_tree_basics() string {
	resString := ""
	arr := []int{43, 1, 123, 13, 324, 46, 75, 2, 265, 500}
	bs := BST{nil}
	bs.Construct(arr)
	res := bs.InOrder()
	resString += strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), ","), "[]") + "\n"
	bs.DeleteItem(13)
	res = bs.InOrder()
	resString += strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), ","), "[]") + "\n"
	bs.DeleteItem(123)
	res = bs.InOrder()
	resString += fmt.Sprint(res)
	return resString
}

type BST struct {
	root *TreeNode
}

func GetBST() *BST {
	return &BST{root:nil}
}

func (bst *BST) Construct(arr []int) {
	for _, val := range arr {
		bst.Insert(val)
	}
}

func (bst *BST) Insert(val int) {
	bst.root = insert(bst.root, val)
}

func insert(node *TreeNode, input int) *TreeNode {
	// if no item, then create one with the input value
	if node == nil {
		return &TreeNode{Val:input, Left:nil, Right:nil}
	}
	if input < node.Val {
		node.Left = insert(node.Left, input)
	} else {
		node.Right = insert(node.Right, input)
	}
	return node
}

func (bst *BST) DeleteItem(val int) {
	deleteNode(bst.root, val)
}

func deleteNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return node
	}
	if node.Val == val {
		// there can be three categories
		// case a) if leave node
		if node.Left == nil && node.Right == nil {
			return nil
		}
		// case b) have just one child
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else { // case c) have both child
			node.Val = getMinAssumingBST(node.Right)
			// delete node with node.Val in the right subtree, note that here the actual deletion does not happen
			// the deletion will happen from the right child of node as we have assign the inorder successor
			node.Right = deleteNode(node.Right, node.Val)
		}
	} else if val < node.Val {
		// try searching in left
		node.Left = deleteNode(node.Left, val)
	} else {
		node.Right = deleteNode(node.Right, val)
	}
	return node
}

func (bst *BST) InOrder() []int {
	arr := make([]int, 0)
	inorder(bst.root, &arr)
	return arr
}

func inorder(node *TreeNode, arrPtr *[]int) {
	if node != nil {
		inorder(node.Left, arrPtr)
		*arrPtr = append(*arrPtr, node.Val)
		inorder(node.Right, arrPtr)
	}
}
