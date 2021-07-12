package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 *
NOTE END */


func Recursion_print_decimal_01(n int) string {

	res := make([]string, 0)
	saveDecimal(n, &res, "")
	return strings.Join(res, "\n")
}

func saveDecimal(n int, ptr *[]string, res string) {

	if len(res) == n {
		if isStingLengthMatchingN(n, res) {
			*ptr = append(*ptr, res)
		}
		return
	}

	for i := 0; i < 10; i++ {
		saveDecimal(n, ptr, res + strconv.Itoa(i))
	}
}

func isStingLengthMatchingN(n int, s string) bool {
	isNonZeroDigitFound := false
	for _, v := range s {
		if !isNonZeroDigitFound && string(v) == "0" {
			return false
		} else {
			isNonZeroDigitFound = true
		}
	}
	return true
}
