package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * As discussed we need to reduce the searching complexity of z-p1, which can also be done using map
 * We will need an extra storage O(n) for the map, where we can store each element in the array
 * Once we start iterating over the input array, we can search for z-p1 in the map in constant time.
 * The time complexity with then be O(n) only
 * Note that if we store all the element in map initially, then we will have duplicate set of items in result
 * so we will insert into map at the time of searching itself.
NOTE END */

func Array_pairs_sum_z_04(arr []int, num int) string {

	res := make([]string, 0)

	m := make(map[int]bool)

	for _, e := range arr {
		if val, ok := m[num-e]; ok && val {
			res = append(res, "(" + strconv.Itoa(e) + "," + strconv.Itoa(num-e) + ")")
		}
		m[e] = true
	}
	return strings.Join(res, ",")
}
