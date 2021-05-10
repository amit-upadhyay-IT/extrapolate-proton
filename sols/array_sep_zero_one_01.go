package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * We can use mechanism used in counting sort, where we can keep count of 0 and 1 and form the array based on the count
 * We can also use the partitioning algorithm used in quick sort, where we will partition the array based on pivot
 * we will keep the pivot element as 1, after the rearrangement we will have all elements lesser than 1, to the left
 * of 1, which is what our desired output.
 * Let us implement that algorithm only here
NOTE END */


func Array_sep_zero_one_01(arr []int) string {
	rearrangeElements(arr)
	res := make([]string, 0)
	for _, e := range arr {
		res = append(res, strconv.Itoa(e))
	}
	return strings.Join(res, ",")
}

func rearrangeElements(arr []int) {
	pivot := 1
	partitionIndex := 0
	leftPtr := 0

	for i := leftPtr; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			swap2(arr, i, partitionIndex)
			partitionIndex++
		}
	}
	arr[partitionIndex] = 1
}

func swap2(arr []int, ind1, ind2 int) {
	temp := arr[ind1]
	arr[ind1] = arr[ind2]
	arr[ind2] = temp
}
