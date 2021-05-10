package utils

import "strconv"

func IntegerSliceToStringSlice(arr []int) []string {
	var res []string
	for _, e := range arr {
		res = append(res, strconv.Itoa(e))
	}
	return res
}
