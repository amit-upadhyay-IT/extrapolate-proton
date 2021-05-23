package sols

import "strconv"

/* NOTE BEGIN
 * Recursion: Defining an opration in terms of itself. With recursion you can solve the problem by solving the smaller
 * occurrences of the problem. Recursions can help solve certain kind of problems very easily.

 * Here the base case would be the one whose answer you already know. We know the 0th power of any number is one.
 * The recursive case would describing the problem interms of itself, i.e. asking the power of the number one less than
 * the current number and multiplying that with the current number.
 * The recurrence relation would look like: pow(n) = n*pow(n-1)

 * Let’s try to see the call stack made interms of a tree:

				  6^7	-----------> returns 6*46656 to main function
				  /
				6^6		-----------> returns 6*7776
				/
			  6^5		-----------> returns 6*1296
			  /
			6^4		-----------> returns 6*216
			/
		  6^3		-----------> returns 6*36
		  /
		6^2			-----------> return 6*6
		/
	  6^1			-----------> return 6*1
	  /
	6^0 			-----------> returns 1 to the parent node

NOTE END */


func Recursion_power_function_01(x, y int) string {
	return strconv.Itoa(power(x, y))
}

// x^y: i.e. multiplying x, y number of times
func power(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}
