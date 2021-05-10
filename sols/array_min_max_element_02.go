package sols

import (
	"math"
	"strconv"
)

/* NOTE BEGIN
 * If we combine elements in a pair and then try to find out the maximum and minimum among the pair only, then
 * how many comaprision will it take?
 * eg: [4, 1]
 * if a[i] > a[i+1] max = a[i]; min = ar[i+1]
 * we can decide it in one comparision only.
 * but, now we need to decide if the maxium of minium in the current pair is maximum or minimum amoung whole array or
 * not. So, for this, we need two more comparision, eg:
 * if localMaximum > globalMaximum : globalMaximum = localMaximum and same for minimum too.
 * i.e. for each two elements in the array we are doing overall 3 comparision
 * So, overall comparision will be
 * C(2) = 3;
 * C(3) = C(2) + 2
 * C(4) = 6;
 * C(5) = C(4) + 2
 * C(n) = (3/2)*n + c; if n is even
 * 		= C(n-1) + c ; c is 2, n is odd
 * This is better than 2n comparision
NOTE END */


func Array_min_max_element_02(arr []int) string {


	minE, maxE := math.MaxInt32, math.MinInt32

	var i int
	if len(arr)%2 == 1 {
		if arr[0] < minE {
			minE = arr[0]
		}
		if arr[0] > maxE {
			maxE = arr[0]
		}
		i = 1
	}

	for i < len(arr) {
		localMin, localMax := arr[i], arr[i+1]
		if arr[i] > arr[i+1] {
			localMin = arr[i+1]
			localMax = arr[i]
		} else {
			localMin = arr[i]
			localMax = arr[i+1]
		}

		if localMin < minE {
			minE = localMin
		}
		if localMax > maxE {
			maxE = localMax
		}
		i += 2
	}

	return strconv.Itoa(minE) + ", " + strconv.Itoa(maxE)

}
