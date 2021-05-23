package sols

import (
	"sort"
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * We can also use map to store the elements in the array for improving the searching of x-a-b, note that a and b are
 * still required to be found out by comparing every possibility of two elements in the array.
 * This will reduce the time complexity to O(n^2) but it will take space of O(n). This is still a better solution
 * O(n^2logn), however in practical cases this solution isn't far better than the previous solution as logn isn't
 * a big number for very large inputs and this one is also taking auxiliary memory O(n). Anyway I have extra time :P
 * let me implement this one too.
NOTE END */



func Array_triplets_for_sum_x_02(arr []int, x int) string {

	elementsMap := make(map[int]int)

	// in map I am storing the key are the element and the value is the index, reason is I will search in map on basis
	// of index as well, so that we avoid searching duplicate elements
	for i, e := range arr {
		elementsMap[e] = i
	}

	tuple3List := make([]Tuple3_1, 0)
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if elementPresentInMap(elementsMap, x-arr[i]-arr[j], i, j) {
				tuple3List = append(tuple3List, Tuple3_1{a: arr[i], b: arr[j], c:x-arr[i]-arr[j]})
			}
		}
	}
	res := make([]string, 0)
	uniqueTup := getUniqueTuples(tuple3List)
	for _, t := range uniqueTup {
		res = append(res, "(" + strconv.Itoa(t.a) + ", " + strconv.Itoa(t.b) + ", " + strconv.Itoa(t.c) + ")")
	}
	return strings.Join(res, ",")
}

func elementPresentInMap(elementsMap map[int]int, key int, ind1, ind2 int) bool {
	if v, ok := elementsMap[key]; ok && (v != ind1 && v != ind2) {
		// (v != ind1 && v != ind2) condition avoids searching same element i.e. element at index ind1 and ind2 are
		// currently being considered as two elements in the possible triplets, i.e. we should not ideally consider them
		return true
	}
	return false
}


type Tuple3_1 struct {
	a,b,c int
}

// returns the unique tuples only, i.e. Tuple3{1,2,3} and Tuple3{2,1,3} will be considered same
func getUniqueTuples(tupleList []Tuple3_1) []Tuple3_1 {

	res := make([]Tuple3_1, 0)

	m := make(map[string]bool)
	for _, tup := range tupleList {
		a := tup.a
		b := tup.b
		c := tup.c

		arr := []int{a, b, c}
		sort.Ints(arr)
		key := strconv.Itoa(arr[0]) + "|" + strconv.Itoa(arr[1]) + "|" + strconv.Itoa(arr[2])

		if _, ok := m[key]; !ok {
			// this key isn't present, i.e. this is not duplicate
			m[key] = true
			res = append(res, tup)
		}
	}
	return res
}
