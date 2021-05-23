package sols

import (
	"sort"
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * If sort the array and then iterate through it, we would be able to keep the count of each element in one go (as after
 * sorting the same elements will be next to each other). Then only we can decide if they are occurring more than n/k
 * time or not.
NOTE END */


func Array_find_element_occur_more_nbyk_unsorted_02(arr []int, k int) string {
	sort.Ints(arr)

	res := make([]string, 0)
	elementCounter := 1
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			elementCounter++
		} else {
			elementCounter = 1
		}
		if elementCounter > len(arr)/k {
			res = append(res, strconv.Itoa(arr[i]))
		}
	}
	return strings.Join(res, ",")
}
