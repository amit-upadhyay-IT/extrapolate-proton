package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * We need to find out how many elements in the array occur more than n/k times in the array
 * We could simply iterate over the input array and keep count of each element in a hashmap and iterate over the map
 * to decide if that element is occurring more than n/k times.
 * Time: O(n), Space: O(n)
NOTE END */


func Array_find_element_occur_more_nbyk_unsorted_01(arr []int, k int) string {

	dic := make(map[int]int)
	for _, e := range arr {
		dic[e]++
	}

	res := make([]string, 0)
	for key, val := range dic {
		if val > len(arr)/k {
			res = append(res, strconv.Itoa(key))
		}
	}
	return strings.Join(res, ",")
}
