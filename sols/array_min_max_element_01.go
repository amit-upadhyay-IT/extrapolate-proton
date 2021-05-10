package sols

import (
	"math"
	"strconv"
)

/* NOTE BEGIN
 * We can use linear search, we need to count the comparision here
 * if we exclude loop comparision, then we will be doing 2n comparision
NOTE END */


func Array_min_max_element_01(arr []int) string {

	minE, maxE := math.MaxInt32,math.MinInt32

	for i := 0 ; i < len(arr); i++ {
		if arr[i] < minE {
			minE = arr[i]
		}
		if arr[i] > maxE {
			maxE = arr[i]
		}
	}

	return strconv.Itoa(minE) + ", " + strconv.Itoa(maxE)
}



