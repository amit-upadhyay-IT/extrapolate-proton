package sols

import "strconv"

/* NOTE BEGIN
 * We can optimize our linear approach solution by considering this fact that to calculate 6^4, you don’t need to
 * calculate all 6^3, 6^2, 6^1 and 6^0. You can just calculate 6^2 and multiply it with itself i.e. 6^2.

 * Here is recursive cases will get split into two parts:
 * When the exponent is odd.
 * When the exponent is even.


 * Let’s look at the recursion stack using a binary tree (as two branches can be formed).

                      6^21
                      /
                     6^20
                  /       \
               6^10       6^10
              /    \      /  \
            6^5    6^5   6^5  6^5
            /       /   /     /
          6^4      6^4 ............
          /  \
        6^2  6^2 ............
        / \
      6^1  6^1 ............

 * The time complexity will still be the same i.e. O(n) because we are still computing the same number of opreations,
 * but the recursion stack space will get reduces to O(log2(n)).

NOTE END */


func Recursion_power_function_02(x, y int) string {

	return strconv.Itoa(power2(x, y))
}


func power2(x, y int) int {
	if y == 0 {
		return 1
	}

	if y % 2 == 0 {  // exponent is even, i.e. we can calculate faster
		return power2(x, y/2) * power2(x, y/2)
	} else {
		return x * power2(x, y-1)
	}
}
