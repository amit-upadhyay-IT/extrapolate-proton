package sols

import (
	"math"
	"sort"
	"strconv"
)

/* NOTE BEGIN
 * We can that smallest positive number would start from 1 (or 0), so we may need to search for numbers starting from 1
 * in the input array. Searching for one element would take O(n) time if we perform a linear search.
 * In worst case scenario we may need to search till n (len of input array) coz once we exceed the search till n, we
 * would simply say n+1 is the answer, right? i.e. n+1 is the smallest positive number not available in the input
 * array. So, this approach would take O(n^2) time.

 * We can improve this time complexity by not doing linear search, we can either sort the array and do binary search
 * Or we can use map to store array elements and then do the search operation on constant time.
 * Let me implement the solution where I sort the array. Once array is sorted, we can simply iterate over that array
 * once to decide the answer, while iteration we need to check if iteration index is matching the array element or not.
 * If there are negative elements in the list, then ignore those elements
 * Time complexity would be O(nlogn)+O(n).
NOTE END */


func Array_smallest_positive_missing_number_01(arr []int) string {

	sort.Ints(arr)

	positiveElementCtr := 1
	prevElement := -math.MaxInt32
	for _, v := range arr {
		if v <= 0 {
			continue
		}
		if positiveElementCtr < v {
			return strconv.Itoa(positiveElementCtr)
		}
		if prevElement != v {
			positiveElementCtr++
		}
		prevElement = v
	}

	return strconv.Itoa(positiveElementCtr)
}
