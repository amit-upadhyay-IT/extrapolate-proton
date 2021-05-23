package sols

import (
	"strconv"
)

/* NOTE BEGIN
 * How will se solve this problem in O(n) time without using extra space?
 * Basically with the use of map we were looking up if there is any such element present in the array or not, right?
 * And we already know that if size of input array is n, then the possible answer would always be less than or equal to
 * n+1, right?
 * So why not use array itself to search, we can make the array indices and mark the element at the index as negative.
 * Eg: input array : [2, 3, -7, 6, 8, 1, -10, 15]
 * We can go to index 2 and mark the element as negative.
 * Next, go to index 3 and mark element as negative, and so on. For searching, we just need to go to that index and see
 * if value present is negative or not. If it's not negative, then we will come to the conclusion that it's the smallest
 * positive number not available in input array.
 * This approach will work for non-negative elements only, for that we can remove the non-negative elements from the
 * array.
NOTE END */


func Array_smallest_positive_missing_number_03(arr []int) string {

	partitionIndex := rearrangeArrayOnElement(arr)
	arr = arr[partitionIndex:]
	for _, e := range arr {
		if getIntMod2(e)-1 < len(arr) {
			arr[getIntMod2(e)-1] = -1 * getIntMod2(arr[getIntMod2(e)-1])
		}
	}

	// run a loop from 1 to n and search for each of the items in the input array
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			// we didn't find i in input array, so this should be the answer
			return strconv.Itoa(i+1)
		}
	}
	return strconv.Itoa(len(arr)+1)
}

func getIntMod2(n int) int {
	if n < 0 {
		return n*-1
	}
	return n
}

func rearrangeArrayOnElement(arr []int) int {

	partiotionIndex := 0
	pivot := 1  // we will partition based on this pivot, i.e. all elements greater than 1 will be to the right of it
	begInd := 0
	endInd := len(arr)-1

	for i := begInd; i <= endInd; i++ {
		if arr[i] < pivot {
			swap3(arr, partiotionIndex, i)
			partiotionIndex++
		}
	}
	return partiotionIndex

}

func swap3(arr []int, ind1, ind2 int) {
	temp := arr[ind1]
	arr[ind1] = arr[ind2]
	arr[ind2] = temp
}

