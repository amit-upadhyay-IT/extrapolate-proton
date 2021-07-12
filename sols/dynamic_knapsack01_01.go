package sols

import "strconv"

/* NOTE BEGIN
 * I am gonna write some introduction first
 * Given a knapsack maximize the profit if we are given with object weight and each object profit. Here either object
 * will be included or excluded, we can not add fraction of object.

 * Note that if we could add fraction of object we will then find out the ratio of weight and profit, then we will
 * we will sort it out in descending order then we will start adding those items with maximum ration till the point the
 * bag is filled completely.
 * But, we are not allowed to add fraction of objects then we might end of leaving some space in the end and we would
 * not be sure if the current items in the knapsack will yeild maximum profit.

 * An example would be:
 * Lets say capacity of knapsack is `C` and there are 2 objects available.
 *         object1   object2
 * profit : 2         C
 * weight : 1         C
 * p/w    : 2         1
 *
 * Now, our greedy method will choose the object1 first coz, p/w ratio is greater, so it will put object1 in the
 * knapsack, the remaining capacity of knapsack would be C-1 and current profit would be 2. Now, we can not put the
 * next item as its weight is C, but we are only left with C-1 capacity

 * Now, lets see if we can solve this using dynamic programming. For that we need to know if there is any sub-structure
 * in the problem and we need to confirm that the problems are repeating.

 * For each item we will have two choices either we can include it or we do not include it. If we include the item, our
 * next result calculation will be different and if we do not include that item, our result calculation will be
 * different. And our target is to maximize the profit.

 * KS(n, C) = max(pn + KS(n-1, C-wi), KS(n-1, C))
 * 			= KS(n-1, C); C < wi, i.e. capacity available is lesser then weight of current item, we will not be able to include it
 * 			= 0 ; n = 0 or C = 0; i.e. no more items are present or capacity is full

 * The above equation is our recursion substructure equation, in this solution we will try to implement this recursive

 * Let's see the recursive tree, a binary tree will get formed coz of double recursion

 * Input: {"weight":[1,2,4],"profit":[10,12,28],"capacity":6}

 *                         KS(3, 6)-0
 *                       /              \
 *                      /                \
 *                KS(2,5)-10              KS(2, 6) - 0
 *               /          \             /            \
 *              /           \            KS(1,4)-12     KS(1,6)-0
 *             /            \           /        \             /   \
 *            /             \        KS(0,0)-40  KS(0,4)-12   /     \
 *          KS(1, 3)-22     KS(1,5)-10                    KS(0,2)-28|KS(0,6)-0
 *          /       \       /          \
 *         /         \     KS(0,1)-38   KS(0,5)-10
 *    KS(0,3)-22   KS(0,3)-22
 *
 * The number of nodes in the above tree at max could be O(2^n) and for computation of each not time taken is constant.
 * So, time complexity of this solution is O(2^n)
 *
NOTE END */


func Dynamic_knapsack01_01(weight, profit []int, capacity int) string {

	maxProfit := recKnapsack(profit, weight, capacity, 0)

	return strconv.Itoa(maxProfit)
}


func recKnapsack(profit, weight []int, cap, itemNum int) int {

	// base condition
	if cap == 0 || itemNum == len(profit) {
		return 0  // there is no capacity or all items are put in knapsack so there can be no profit
	}

	if cap < weight[itemNum] {
		return recKnapsack(profit, weight, cap, itemNum+1)  // we can not include the item is capacity available is lesser than current items weight
	}

	profIncCurr := profit[itemNum]+recKnapsack(profit, weight, cap - weight[itemNum], itemNum+1)
	profWithoutCurr := recKnapsack(profit, weight, cap, itemNum+1)
	return getMaxElement(profIncCurr, profWithoutCurr)
}

func getMaxElement(a, b int) int {
	if a > b {
		return a
	}
	return b
}
