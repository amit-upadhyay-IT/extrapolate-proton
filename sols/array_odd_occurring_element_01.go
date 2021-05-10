package sols

import "strconv"

/* NOTE BEGIN
 * A simple approach could be for every element in array, you can count total occurrence of that element in the entire
 * array, but this would cost us O(n^2)
 * Another approach could be using a map to count the number of occurrence of elements while iterating over array, and
 * this would only cost us O(n) in time and space both.
 * Using a very simple bitwise operation we can solve this problem in O(n) time and without extra space.
 * The xor operator can help us here as a xor a = 0 and a xor a xor a = a
NOTE END */

func Array_odd_occurring_element_01(arr []int) string {
	var res int
	for _, e := range arr {
		res ^= e
	}
	return strconv.Itoa(res)
}
