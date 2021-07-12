package sols

import (
	"strconv"
)

/* NOTE BEGIN
 * In the previous solution we can have lots of common computations happening again and again. But if we observe
 * the previous solution the maximum unique sub-problems we could ever have is N*C, where N is number of items and C is
 * the capacity.
 * Why is this N*C?
 * In the recursive equation where input is KS(N,C), for each value of N, i.e. for each item we many need to consider
 * C capacity where C can vary from 1, 2, 3...C. i.e. in worst case we will have to compute N*C problems.
 *
 * Now, if we have C < 2^N, then we can improve the recursive solution by storing the previously computed values
 * in a data structure. And in this case our solution time could further be reduced.

 * Basically, we will create a 2D matrix of size N+1 * C+1, and there I will store the computed values. Once if a cell
 * is computed and whenever the same sub problem is occurring again, I will use the values stored in my solution data
 * structure. And the reason we need to allocate one extra row and column is to store the base case.
 * Why do we want to store the base case? to make our solution generic. You will understand once you see the code.
 * to see the base condition, please check the previous solution recursive equation.
 * This algorithm is also called a dynamic programming algorithm, DP can be applied in bottom up fashion or top down
 * fashion.
 * Currently, we will do a bottom up dynamic programming solution, it is bottom up because, we first compute the
 * smallest problem and using those computation, I started building up the bigger values

 * Input: {"weight":[1,2,4],"profit":[10,12,28],"capacity":6}

 *      0     1     2     3     4     5     6
 *   |     |  0  |  0  |  0  |  0  |  0  |  0  |
 * 1 |  0  |  10 |  10 |  10 |  10 |  10 |  10 |
 * 2 |  0  |  10 |  12 |     |     |     |     |
 * 3 |  0  |     |     |     |     |     |     |
NOTE END */

func Dynamic_knapsack01_02(weight, profit []int, capacity int) string {

	return strconv.Itoa(dpKnapsack(profit, weight, capacity))
}

func dpKnapsack(profit, weight []int, capacity int) int {

	// initialize the solution data structure, note that all cells are initialized with zero
	sol := make([][]int, len(profit)+1)
	for i := range sol {
		sol[i] = make([]int, capacity+1)
	}

	// KS(n, C) = max(pi+KS(n-1, C-ci), KS(n-1, C))
	//          = if C < ci; KS(n-1, C)

	for i := 1; i <= len(profit); i++ {

		for j := 1; j <= capacity; j++ {
			// if the current capacity of knapsack is lesser than item capacity, then do not include
			if j < weight[i-1] {
				sol[i][j] = sol[i-1][j]
			} else {
				profit1 := profit[i-1]+sol[i-1][j-weight[i-1]]
				profit2 := sol[i-1][j]
				sol[i][j] = getMaxDP2(profit1, profit2)
			}
		}
	}

	return sol[len(profit)][capacity]
}

func getMaxDP2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
