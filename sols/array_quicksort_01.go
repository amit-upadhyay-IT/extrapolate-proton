package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * This is a recursive sorting technique, where we change position of one element (also called as pivot) in each
 * recursion call, by position change I mean, we place the pivot element at it's appropriate position in the array.
 * Appropriate position means, all element greater than pivot are placed on right of pivot element and all elements
 * lesser than pivot are placed on left of pivot position(in case of ascending sort).
 * The selection of pivot element in the array doesn't matter, as the main concern is placing one element at it's
 * appropriate position.
 * I will tell about this process of placing pivot at its appropriate position later, but for now if we are able to
 * do so recursively for each sub parts of the array, we would finally be able to sort the array, right?
 * The rearrangement process:
 * 		- We can iterative over array, and for each element we see while iterating we need to swap the elements position
 * 		  such that at the end of iteration we should only see elements smaller than pivot to the left of array.
 * 		- For this, we should always make sure that elements smaller than pivot should never be at right side of pivot
 * 		  i.e. if we find any element smaller than pivot we should swap the element with some other element, right?
 * 		- How to decide which element to swap with?
 * 		- To decide which element to swap with, we can maintain a counter variable (or a pointer variable, which will
 * 		  point at array indices) which can tell us how many elements are smaller than pivot till the current point of
 *		  iteration, and we will swap the element at `i` with element at this counter variable.
 * So, basically, we partition the array based on pivot element and then again apply same rearrangement logic on the
 * two subarray, i.e. the problem is split into two parts, right? Similarly, we need to solve it for their sub parts
 * and so on, untill the sub array has size of 1.
 *
 * Time complexity: T(n) = 2T(n/2) + c*n ; if n > 1; c is some constant
 * 						 = c 			 ; if n = 1
 * i.e. O(nlogn)
 * Note that this process of sorting will mutate the passed array.
NOTE END */


func Array_quicksort_01(arr []int) string {
	quickSort(arr, 0, len(arr)-1)
	return strings.Join(intSliceToStringSlice(arr), ",")
}


func quickSort(arr []int, firstInd, lastInd int) {

	// we need to recursively sort the array inplace until there is just one element left in the smallest sub array
	// to know if there is just one element left in the subarray, we can compare the firstIndex and lastIndex values
	if firstInd < lastInd {

		partitionIndex := rearrangeArrayOnPivot(arr, firstInd, lastInd, lastInd) // keeping the pivot element as the last one always in the array
		quickSort(arr, firstInd, partitionIndex-1)
		quickSort(arr, partitionIndex+1, lastInd)
	}
}


// this method places the pivot element such that all elements to lesser than it are on left of it and elements
// greater than it are right of it, it will also return the partition index, i.e. the index where pivot is placed
// after rearrangement
func rearrangeArrayOnPivot(arr []int, begInd, endInd, pivotIndex int) int {

	// we need to maintain a counter variable which will tell us about the count of elements which are smaller
	// than pivot element, this will help us decide the swapping element whenever we encounter any element
	// smaller than the pivot.
	partitionIndex := begInd

	for i := begInd; i < endInd; i++ {

		if arr[i] < arr[pivotIndex] { // but why? because i is always ahead or equal to partitionIndex and we want
			// no smaller elements than pivot should be after partitionIndex thus we swap these
			swap(arr, i, partitionIndex)
			partitionIndex++
		}
	}
	swap(arr, partitionIndex, endInd)

	return partitionIndex
}

func swap(arr []int, ind1, ind2 int) {
	temp := arr[ind1]
	arr[ind1] = arr[ind2]
	arr[ind2] = temp
}

func intSliceToStringSlice(arr []int) []string {
	res := make([]string, 0)
	for _, e := range arr {
		res = append(res, strconv.Itoa(e))
	}
	return res
}
