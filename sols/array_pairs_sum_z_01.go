package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * find all possible pairs from an array, whose sum is equal to given input in the program
 * We can use brute force approach, for each element p1 search for z-p1 in array
NOTE END */

func Array_pairs_sum_z_01(arr []int, z int) string {
	// iterate over arr, for each element p1, search into array for p2 whose sum is equal to z

	res := make([]string, 0)
	for i, p1 := range arr {
		for _, p2 := range arr[i+1:] {
			if p1 + p2 == z {
				res = append(res, "(" + strconv.Itoa(p1) + "," + strconv.Itoa(p2) + ")")
			}
		}
	}
	return strings.Join(res, ",")
}
