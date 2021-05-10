package sols

import (
	"sort"
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * find all possible pairs from an array, whose sum is equal to given input in the program
 * in the last solution, we were assuming one element as the first pair and then we were searching for the
 * second element by iterating over the array.
 * We can make our search operation faster, either by sorting the array or by using a map.
 * let's see at the solution where we will sort the array.
 * Solution time complexity: O(nlogn) for sorting and O(logn) for searching, i.e. overall it will be O(nlogn)
NOTE END */

func Array_pairs_sum_z_02(arr []int, z int) string {

	res := make([]string, 0)
	sort.Ints(arr)
	for _, e := range arr {
		if binarySearch(arr, z-e) {
			res = append(res, "(" + strconv.Itoa(e) + "," + strconv.Itoa(z-e) + ")")
		}
	}
	// in this approach we will have duplicate pairs, so removing the duplicate entries from res
	res = res[0:len(res)/2]
	return strings.Join(res, ",")
}


func binarySearch(arr []int, key int) bool {
	// either key is in left part of arr or right part of array or its the middle element of array
	arrLen := len(arr)
	if len(arr) == 0 {
		return false
	}
	midElement := arr[arrLen/2]
	if midElement == key {
		return true
	}
	if key < midElement {
		return binarySearch(arr[0:arrLen/2], key)
	} else {
		return binarySearch(arr[(arrLen/2)+1: arrLen], key)
	}
}
