package sols

import "strconv"

/* NOTE BEGIN
 * to find the maximum different between two elements s.t larger element is on right of smaller element
 * We can use brute force, i.e. consider one element as pair one and then search for the next biggest element in the
 * entire set.
 * This will hold a time complexity of O(n^2)j
NOTE END */


func Array_max_diff_two_elements_larger_after_smaller_01(arr []int) string {
	maxDiff := 0

	for i := len(arr)-1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if maxDiff < arr[i] - arr[j] {
				maxDiff = arr[i] - arr[j]
			}
		}
	}
	return strconv.Itoa(maxDiff)
}
