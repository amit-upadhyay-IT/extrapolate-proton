package sols

import (
	"strconv"
)

/* NOTE BEGIN
 * I will try to write a top down approach for this solution
NOTE END */


func Dynamic_knapsack01_03(weight, profit []int, capacity int) string {

	sol := make([][]int, len(profit))
	for i := range sol {
		sol[i] = make([]int, capacity)
		for j := 0; j < capacity; j++ {
			sol[i][j] = -1
		}
	}

	p := recKnapsack2(profit, weight, capacity, 0, sol)
	return strconv.Itoa(p)
}


// KS(i, C) = max(pi + KS(i-1, C-ci), KS(i-1, C))
func recKnapsack2(profit, weight []int, capacity int, itemNum int, sol [][]int) int {

	if capacity == 0 || itemNum == len(profit) {
		return 0
	}

	if sol[itemNum][capacity-1] != -1 {
		return sol[itemNum][capacity-1]
	}

	if capacity < weight[itemNum] {
		p := recKnapsack2(profit, weight, capacity, itemNum+1, sol)
		sol[itemNum][capacity-1] = p
		return p
	}

	prof1 := profit[itemNum] + recKnapsack2(profit, weight, capacity-weight[itemNum], itemNum+1, sol)
	prof2 := recKnapsack2(profit, weight, capacity, itemNum+1, sol)

	if prof1 > prof2 {
		sol[itemNum][capacity-1] = prof1
		return prof1
	} else {
		sol[itemNum][capacity-1] = prof2
		return prof2
	}

}
