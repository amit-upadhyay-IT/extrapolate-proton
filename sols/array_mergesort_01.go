package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * This is also a recursive sorting technique, where we divide the array into sub arrays and try to sort the sub arrays
 * After the subarrays are sorted, we merge the sub arrays into the final array, which will result in sorted array.
 * Detailed explanation:
 * Here we break the problem into sub problems
 * We divide the array into two possible equal half, we find the middle element and divide the array into parts where
 * first part is all the elements to the left of middle element and second part is all elements to right of middle
 * element.
 * Now, if we are able to sort the two half, then we can merge them together and that's all. Very simple right?
NOTE END */

func Array_mergesort_01(arr []int) string {

	mergeSort(arr)
	return strings.Join(intSliceToStringSlice2(arr), ",")
}


func mergeSort(arr []int) {

	// base condition? one element array is always sorted, so return that
	if len(arr) < 2 {
		return
	}

	midInd := len(arr)/2
	// allocate two array and assign values to it
	leftSubArr := make([]int, midInd)
	rightSubArr := make([]int, len(arr)-midInd)

	for i := 0; i < midInd; i++ {
		leftSubArr[i] = arr[i]
	}
	for i := midInd; i < len(arr); i++ {
		rightSubArr[i-midInd] = arr[i]
	}

	mergeSort(leftSubArr)
	mergeSort(rightSubArr)
	mergeArrays(arr, leftSubArr, rightSubArr)
	//fmt.Print(arr); fmt.Print(leftSubArr); fmt.Print(rightSubArr)
	//fmt.Println()
}

func mergeArrays(masterArr, arr1, arr2 []int) {
	ctr1, ctr2, mstrCtr := 0, 0, 0

	for ctr1 < len(arr1) && ctr2 < len(arr2) {
		if arr1[ctr1] < arr2[ctr2] {
			masterArr[mstrCtr] = arr1[ctr1]
			ctr1++
		} else {
			masterArr[mstrCtr] = arr2[ctr2]
			ctr2++
		}
		mstrCtr++
	}

	for ctr1 < len(arr1) {
		masterArr[mstrCtr] = arr1[ctr1]
		mstrCtr++; ctr1++
	}

	for ctr2 < len(arr2) {
		masterArr[mstrCtr] = arr2[ctr2]
		mstrCtr++; ctr2++
	}
}

func intSliceToStringSlice2(arr []int) []string {
	res := make([]string, 0)
	for _, e := range arr {
		res = append(res, strconv.Itoa(e))
	}
	return res
}