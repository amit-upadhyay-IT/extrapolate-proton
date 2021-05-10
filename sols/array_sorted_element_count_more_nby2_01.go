package sols

import "strconv"

/* NOTE BEGIN
 * since array is already sorted, if we find any occurrence where arr[i] is same as arr[i+n/2], then we can say that
 * there is a element in array which is occurring more than n/2.
NOTE END */


func Array_sorted_element_count_more_nby2_01(arr []int) string {

	var res string = "NaN"

	for i, _ := range arr[:len(arr)/2] {
		if arr[i] == arr[i+len(arr)/2] {
			res = strconv.Itoa(arr[i])
		}
	}

	return res

}
