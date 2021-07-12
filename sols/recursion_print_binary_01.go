package sols

import (
	"strings"
)

/* NOTE BEGIN
 * Do you see how is this problem self-similar?
 * To print n digit long binary number, I want all the preceding number of binary number and prepend 0 or prepend 1
 * to it.
NOTE END */


func Recursion_print_binary_01(n int) string {
	res := make([][]string, 0)
	temp := make([]string, 0)
	ptrToRes := &res
	saveBinary(n, temp, ptrToRes)

	result := ""
	for _, v := range res {
		result += strings.Join(v, ",")
		result += "\n"
	}
	return result
}


func saveBinary(n int, temp []string, ptrToRes *[][]string) {

	// base condition, if result length is n
	if len(temp) == n {
		c := make([]string, len(temp))
		copy(c, temp)  // if we use a immutable type structure like string, we do not need to do deep copy for temp
		*ptrToRes = append(*ptrToRes, c)
		return
	} else {
		temp = append(temp, "0")
		saveBinary(n, temp, ptrToRes)

		temp = temp[:len(temp)-1]
		temp = append(temp, "1")
		saveBinary(n, temp, ptrToRes)
	}
}
