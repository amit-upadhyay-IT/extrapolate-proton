package sols

import (
	"math"
	"sort"
	"strconv"
)

/* NOTE BEGIN
 * We need to find a pair whose sum is closest to zero,
 * eg: [15, 5, -20, 30, -45], (15, -20)
 * One way could be iteratively searching for pairs which would result in O(n^2) time complexity
 * We need optimization here, note that we do not have exact value to which the sum of pair should match, we just have
 * a relative value (depending on input array elements) which should be close to zero.
 * If we use two pointer approach, and at the same time we are able to decide which direction to move the pointer
 * while iterating over the array, then we would be able to find such pairs, right?
 * To know which direction to move the pointers, we need to sort the array first, coz if array is sorted, we would be
 * able to know if sum is negative, then we should try to increase the sum and for that we need to move the left pointer
 * to next position, else if sum is positive we are sure that we need to decrease the pair sum, for which we need
 * to decrease the right pointer value.
 * Time complexity will be **O(nlog)**, this is due to sorting. If input array is already sorted, then complexity will just
 * be O(n)
 * Keeping in mind, let's just start writing code.
NOTE END */


func Array_pair_sum_close_zero_01(arr []int) string {

	sort.Ints(arr)

	leftPtr, rightPtr := 0, len(arr)-1

	minPairSum := math.MaxInt32
	var res string

	for leftPtr < rightPtr {
		pairSum := arr[leftPtr] + arr[rightPtr]
		if pairSum < 0 {
			if getIntMod(pairSum) < minPairSum {
				minPairSum = getIntMod(pairSum)
				res = strconv.Itoa(arr[leftPtr]) + ", " + strconv.Itoa(arr[rightPtr])
			}
			leftPtr++
		} else {
			if getIntMod(pairSum) < minPairSum {
				minPairSum = getIntMod(pairSum)
				res = strconv.Itoa(arr[leftPtr]) + ", " + strconv.Itoa(arr[rightPtr])
			}
			rightPtr--
		}
	}
	return res
}

func getIntMod(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
