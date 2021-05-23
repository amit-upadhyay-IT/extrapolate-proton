package sols

import "strconv"

/* NOTE BEGIN
 * In this solution I will use map to store input array elements, so that searching could be done in constant time
 * Basically, we will start searching from 1 to len(arr), arr is input array. As soon as we don't find the element
 * that would be the answer to the question
NOTE END */

func Array_smallest_positive_missing_number_02(arr []int) string {

	m := make(map[int]bool)
	for _, e := range arr {
		m[e] = true
	}

	ctr := 1
	for i := 1; i < len(arr)+1; i++ {
		if _, ok := m[i]; !ok {
			return strconv.Itoa(i)
		}
		ctr = i
	}

	return strconv.Itoa(ctr)
}
