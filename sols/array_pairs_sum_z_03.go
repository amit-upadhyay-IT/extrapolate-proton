package sols

import (
	"sort"
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * Sort the array and keep two pointers, one on the left most element and another on the right most(biggest element)
 * do sum of element pointed by both the pointers
 * if sum is lesser than z, it means we need to decrease the sum of elements pointed by pointer and we can do so
 * by moving right most pointer to a lower position
 * else if, sum is greater than z, we need to increase the sum of elements pointed by the two pointers and so we
 * move left pointer to the next position
 * else, is sum is equal to z, we add it our result set.
 * Note that we would need to perform this operation until left ptr and right ptr collide
NOTE END */

func Array_pairs_sum_z_03(arr []int, z int) string {

	res := make([]string, 0)

	if arr == nil || len(arr) == 0 {
		return strings.Join(res, ",")
	}
	sort.Ints(arr)
	leftPtr := 0
	rightPtr := len(arr)-1

	for leftPtr < rightPtr {
		if arr[leftPtr] + arr[rightPtr] < z {
			// we need to increase the sum
			leftPtr++
		} else if arr[leftPtr] + arr[rightPtr] > z {
			// we need to decrease the sum
			rightPtr--
		} else {
			res = append(res, "(" + strconv.Itoa(arr[leftPtr]) + "," + strconv.Itoa(arr[rightPtr]) + ")")
			leftPtr++
			rightPtr--
		}
	}
	return strings.Join(res, ",")
}
