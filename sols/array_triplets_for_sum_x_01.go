package sols

import (
	"sort"
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * We need to find a, b and c from array elements so that their sum is x (input given).
 * One simple method would be iterating over array and consider every possible combination, i.e. nC3 possibilities
 * and see if their sum is x or not, and this would take O(n^3) time.
 * I will not implement this, let's think of better solutions.
 * One solution could be sorting the array so that we can improve the searching for (x-a-b) after considering all
 * possible combination of a and b. i.e. this would take O(n^2logn) time. This is still better solution.
NOTE END */

func Array_triplets_for_sum_x_01(arr []int, x int) string {

	sort.Ints(arr)

	res := make([]string, 0)

	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if binarySearch2(arr, x-arr[i]-arr[j], i, j) {
				res = append(res, "(" + strconv.Itoa(arr[i]) + ", " + strconv.Itoa(arr[j]) + ", " + strconv.Itoa(x-arr[i]-arr[j]) + ")")
			}
		}
	}
	return strings.Join(res, ",")
}

func binarySearch2(arr []int, key int, ignoreInd1, ignoreInd2 int) bool {

	// base condition, if array length is 0
	if len(arr) == 0 {
		return false
	}

	midIndex := len(arr)/2
	// ignoring search on these indices as they have already been considered as an element of the triplet
	if midIndex == ignoreInd1 || midIndex == ignoreInd2 {
		return false
	}
	if arr[midIndex] == key {
		return true
	} else if arr[midIndex] < key {
		return binarySearch2(arr[midIndex+1:], key, ignoreInd1, ignoreInd2)
	} else {
		return binarySearch2(arr[:midIndex], key, ignoreInd1, ignoreInd2)
	}

}
