package sols

import (
	"strconv"
)

/* NOTE BEGIN
 * We want to find the solution in O(n) time only or lesser if possible
 * We can use map to store count of each element, and iterating over map we can check if any of the element is occurring
 * more than n/2 times.
 * But space complexity will be O(n) also.
NOTE END */

func Array_find_element_occur_more_nby2_unsorted_01(arr []int) string {

	dic := make(map[int]int)

	for _, e := range arr {
		dic[e]++
	}

	for k, v := range dic {
		if v >= len(arr)/2 {
			return strconv.Itoa(k)
		}
	}
	return "NaN"
}
