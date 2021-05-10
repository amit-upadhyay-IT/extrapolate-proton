package sols

import (
	"strconv"
)

/* NOTE BEGIN
 * in a sorted array if there is any element occurring more than n/2 time then it must be present in the middle of the
 * array, we can simply find the first and last occurrence of mid element in the array and if the difference between
 * their first and last position came out to be greater than n/2 we would say yes, this is it!
 * time complexity will be O(logn)
NOTE END */

func Array_sorted_element_count_more_nby2_02(arr []int) string {

	var res string = "NaN"

	if len(arr) < 1 {
		return res
	}

	firstOccrInd := getOccurrencePosition(arr, arr[len(arr)/2], 0, len(arr)-1,  -1, true)
	lastOccrInd := getOccurrencePosition(arr, arr[len(arr)/2], 0, len(arr)-1, -1, false)
	if firstOccrInd == -1 || lastOccrInd == -1 {
		return res
	}
	if lastOccrInd - firstOccrInd >= len(arr)/2 {
		res = strconv.Itoa(arr[len(arr)/2])
	}
	return res
}

// since array is sorted we can easily apply divide and conquer to know which side of sub array the element is present
func getOccurrencePosition(arr []int, key, firstPos, lastPos, res int, findFirstOccr bool) int {

	if firstPos <= lastPos {
		midIdx := (firstPos+lastPos)/2
		if arr[midIdx] == key {
			res = midIdx
			if findFirstOccr {
				return getOccurrencePosition(arr, key, firstPos, midIdx-1, res, findFirstOccr)
			} else {
				return getOccurrencePosition(arr, key, midIdx+1, lastPos, res, findFirstOccr)
			}
		} else if arr[midIdx] < key {
			// search in right hand side of array in case array is sorted in ascending order
			return getOccurrencePosition(arr, key, midIdx+1, lastPos, res, findFirstOccr)
		} else {
			// arr[midIdx] > key, i.e. search in first half of the array
			return getOccurrencePosition(arr, key, firstPos, midIdx-1, res, findFirstOccr)
		}
	}
	return res
}
