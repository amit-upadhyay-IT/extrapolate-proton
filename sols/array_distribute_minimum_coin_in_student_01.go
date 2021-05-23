package sols

import (
	"strconv"
)

/* NOTE BEGIN
 * Can we solve this problem with recursion? Well, that a very stupid idea, coz unnecessary memory stack will get used.
 * Note that this recursive solution is the implementation of dp only, it's using the previously calculated values to
 * tell the answer of next sub problem. The real fun will be actually solving with native approach using recursion.
 * What say?
NOTE END */


func Array_distribute_minimum_coin_in_student_01(arr []int) string {

	sol := make([]int, 1)
	sol[0] = 0  // initially we have not distributed any coins, so initializing the count with zero

	coinCtr := 1

	fillSol(arr, 0, sol, 1, &coinCtr)

	return strconv.Itoa(sol[0])
}


func fillSol(arr []int, i int, sol []int, leftCoinCnt int, rightCoinCtr *int) {
	// base condition
	if i >= len(arr) {
		return
	}

	leftIdx, rightIdx := i - 1, i + 1
	if i - 1 < 0 {
		leftIdx = i
	}
	if i + 1 >= len(arr) {
		rightIdx = i
	}

	if arr[i] > arr[leftIdx] {
		leftCoinCnt++
	} else {
		leftCoinCnt = 1
	}

	fillSol(arr, i+1, sol, leftCoinCnt, rightCoinCtr)


	if arr[i] > arr[rightIdx] {
		*rightCoinCtr++
	} else {
		*rightCoinCtr = 1
	}

	sol[0] += getMax(leftCoinCnt, *rightCoinCtr)

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
