package sols

import (
	"strconv"
)

/* NOTE BEGIN
 * [5, 12, 7, 4, 12, 12, 5, 12, 12, 12, 2, 12]
 * Let's split the problem into two parts:
 * 	- find the element which is likely to occur most number of times in the array.
 * 	- validate if the above element is actually occurring more than n/2 time or not.
 * How to get element which is occurring most number of times?
 * what if elements in the array can vote for itself, i.e. while iterating through the array if we see an element
 * occurring we can set the current_max_occurring_element as that element as increase its vote count by 1.
 * Note that, there can be only one majority candidate if it's count has to be more than n/2, so our approach would
 * be to find out one possible majority candidate.
 * So, basically we need to maintain two variables while iterating over the array, one to keep the votes count and
 * another to keep the possible majority candidate element.
 * Whenever we see arr[i] is matching the current majority candidate, we will increase the votes count
 * If arr[i] doesn't match we will decrease the vote count. See code below to understand this process.

 * Time complexity: O(n)
NOTE END */


func Array_find_element_occur_more_nby2_unsorted_02(arr []int) string {

	// find element occurring most number of time
	majElement := getMajorityElement(arr)

	cnt := 0
	for _, e := range arr {
		if majElement == e {
			cnt++
		}
	}
	if cnt >= len(arr)/2 {
		return strconv.Itoa(majElement)
	}
	return "NaN"
}

func getMajorityElement(arr []int) int {
	currentElement := arr[0]
	elementVotes := 1

	for i := 1; i < len(arr); i++ {
		if currentElement == arr[i] {
			elementVotes++
		} else {
			// there can be two conditions for element votes, either it can be zero or positive
			if elementVotes == 0 {
				currentElement = arr[i]
				elementVotes++
			} else {
				elementVotes--
			}
		}
	}
	return currentElement
}


