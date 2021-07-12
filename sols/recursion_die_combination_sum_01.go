package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 *
NOTE END */

func Recursion_die_combination_sum_01(n int, sum int) string {

	resList := make([][]string, 0)
	temp := make([]int, 0)  // a temporary slice to store the dice values in the recursive calls
	diceRoll(n, temp, sum, &resList)

	resToReturn := ""
	for _, v := range resList {
		resToReturn += strings.Join(v, ",")
		resToReturn += "\n"
	}
	return resToReturn
}


func diceRoll(n int, res []int, desiredSum int, resPtr *[][]string) {

	// base condition, if there is no dice to roll
	if n == 0 {
		if desiredSum == 0 {
			*resPtr = append(*resPtr, deepCopyAndConvertToString(res))
		}
		return
	} else if desiredSum >= n * 1 && desiredSum <= n * 6 {
		for i := 1; i <= 6; i++ {
			res = append(res, i)
			diceRoll(n-1, res, desiredSum-i, resPtr)
			res = res[:len(res)-1]
		}
	}

}

func deepCopyAndConvertToString(arr []int) []string {

	res := make([]string, 0)
	for _, v := range arr {
		res = append(res, strconv.Itoa(v))
	}
	return res
}
