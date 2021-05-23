package sols

import "strconv"

/* NOTE BEGIN
 * We can reduce the time complexity of previous solution from O(n) to O(log2(n)) by computing the value just once and
 * storing the computed value in the same stack i.e. memoization.
 *
 * Here the value gets computed only once and it can be used again in the same stack at the time of backtracking.
 * So time complexity is O(log2(n)).
NOTE END */

func Recursion_power_function_03(x, y int) string {
	return strconv.Itoa(power3(x, y))
}


func power3(x, y int) int {
	if y == 0 {
		return 1
	}

	computedPower := power3(x, y/2)
	if y % 2 == 0 {
		// power is even
		return computedPower * computedPower
	} else {
		return x * computedPower * computedPower
	}
}
