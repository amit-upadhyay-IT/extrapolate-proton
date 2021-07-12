package tree

import "strconv"

/* NOTE BEGIN
 * Count the possible structures of tree.
 * We will recursively try to solve this problem
 * Possible tree with 3 nodes: 5 possible tree structures
 *        N              N        N              N            N
 *       / \            /         \             /              \
 *      N  N           N           N           N                N
 *                    /            \            \              /
 *                   N             N            N             N

 * If you observe, for N nodes, lets start with a skewed tree, i.e. lets assume all nodes are at the right size of
 * each of the parent node, eg:
 *         N
 *          \
 *           N
 *            .
 *             .
 *              .
 *               N

 * Now, In the above skewed tree, each node could possibly be the root node, and rest of the node could be left child
 * or right child, eg:

 *           N
 *          / \
 *         N  N
 *             \
 *             N
 *              .
 *               .

 * Similarly, we can increase the number of nodes at left and decrease the number of nodes at right child nodes of root
 * node.
 * Also, lets assume, there can be n1 number of possible structures in left child of root node and there can be n2
 * numbers of possible structures in the right child of root node, then total number of possible structures using
 * left and right trees could be n1*n2, coz each of the possible structures at left child could be combined with each
 * of the possible structures of right trees.

 * Now, we can write our recursive equation:
 * count(n) = for all values 1...n
				sum += count(i-1)*count(n-i)
            = 1; if n <= 1
NOTE END */


func Tree_count_of_binary_trees_01(n int) string {

	return strconv.Itoa(countTree(n))
}

func countTree(n int) int {
	if n <= 1 {
		return 1
	}
	leftCnt, rightCnt := 0, 0
	sum := 0
	for i := 1; i <= n; i++ {
		leftCnt = countTree(i-1)
		rightCnt = countTree(n-i)

		sum += leftCnt*rightCnt
	}
	return sum
}


