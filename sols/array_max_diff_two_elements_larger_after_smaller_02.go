package sols

import "strconv"

/* NOTE BEGIN
 * In the previous approach, i.e. brute force approach, we were essentially assuming maximum element and then searching
 * for the minimum element to the left side of it, where this search operation of minimum element was taking O(n) time.
 * And so over all time was getting O(n^2).
 * If we observe the problem, we basically need to maximise the different between two elements, and the bigger element
 * is always on the right side.
 * Eg: let x and y represent the two elements where y appears after x in input array.
 * We want to maximise (y-x), i.e. we need to minimize `x` as much as possible, i.e. it should be the minimum element
 * present on the left side of y.
 * So, for this we can maintain a minimum so far, while iterating the array from the left side, and wherever the
 * `current_element-minium_so_far` is greatest in the complete array iteration process, that would be the answer.
NOTE END */


func Array_max_diff_two_elements_larger_after_smaller_02(arr []int) string {

	if len(arr) < 2 {
		return "NaN"
	}

	maxDiff := 0
	minSoFar := arr[0]

	for _, e := range arr[1:] {
		if e - minSoFar > maxDiff {
			maxDiff = e - minSoFar
		}
		if e < minSoFar {
			minSoFar = e
		}
	}

	return strconv.Itoa(maxDiff)

}
