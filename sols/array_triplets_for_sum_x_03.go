package sols

import (
	"sort"
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * Can we solve this problem in O(n^2) without any auxiliary space?
 * Yes, we can use this pointer approach in sorted array, right?
 * We can keep one element as our possible element among a, b or c. Keeping `a`, we just need to find pair `b` and `c`
 * such that b + c should be equal to x-a, right?
 * We already know if array we sorted, we can use the left pointer and right pointer approach to find such pairs in O(n)
 * time. We will keep left pointer at beginning of array and right pointer at end of array, if desired sum is greater
 * than sum of elements pointed by these two pointers then we will try to decrease the sum, for that we will decrement
 * the right hand side pointer, and so on.
 * Let's  implement this solution too. Time complexity will be O(n^2).
NOTE END */


func Array_triplets_for_sum_x_03(arr []int, x int) string {
	sort.Ints(arr)

	res := make([]string, 0)

	for i := 0; i < len(arr); i++ {
		a := arr[i]
		pairList := getPairsForSum(arr, x-a, i)
		if len(pairList) > 0 {
			for _, p := range pairList {
				b := p.b
				c := p.c
				// just store all a,b and c
				res = append(res, "(" + strconv.Itoa(a) + ", " + strconv.Itoa(b) + ", " + strconv.Itoa(c) + ")")
			}
		}
	}
	return strings.Join(res, ",")
}


type Pair2 struct {
	b int
	c int
}

// i is the current index, we need to ignore this index as per the solution requirement
// coz if we consider this index we will end up considering duplicate indices in our triplet
func getPairsForSum(arr []int, x int, i int) []Pair2 {

	leftPrt, rightPtr := 0, len(arr)-1

	pairList := make([]Pair2, 0)

	for leftPrt < rightPtr {
		if leftPrt == i || rightPtr == i {  // I am not considering the element which is already considered as part of triplet in the running algorithm
			leftPrt++; rightPtr--;
			continue
		}
		if arr[leftPrt] + arr[rightPtr] == x {
			// store it somewhere
			pairList = append(pairList, Pair2{b:arr[leftPrt], c:arr[rightPtr]})
			// we have tried this leftPtr and rightPtr, i.e. we should change their values and try something new next time
			leftPrt++; rightPtr--;
		} else if arr[leftPrt] + arr[rightPtr] > x {
			// we need to decrease the sum, i.e. we will decrement rightPtr
			rightPtr--
		} else {
			leftPrt++
		}
	}
	return pairList
}
