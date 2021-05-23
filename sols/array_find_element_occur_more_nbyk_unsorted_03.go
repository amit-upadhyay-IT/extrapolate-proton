package sols

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * We have seen the voting algorithm, which we have used while finding the element in array occurring more than n/2
 * times. We can also use similar kind of approach.
 * Since, we are required to find all elements who are occurring more than n/k times, so we can first find out all
 * possible elements which can occur more than n/k times and then check the actual count of those elements after
 * iterating over the array again.
 * i.e. the solution can be divided into two parts:
 * 		- find all possible elements which can occur more than n/k times in the array.
 * 		- validate if those elements actually occur more than n/k times.

 * Now, how do we know all possible elements which occur more than n/k times.
 * We know that the possibility of such elements can never be more than k-1 times, just think more on it if you have
 * doubt.
 * So, we can maintain an array for k-1 size (note that in while finding elements occurring more than n/2 time we just
 * maintained a variable, but here we need to maintain an array because of increase in possibility).
 * The above array will count the number of times element is voting for itself, and voting will increase if we encounter
 * same element while iterating over the array, voting will decrease if we find new element (which are not present
 * in our possibility array) while iterating over the input array. If you have got fair idea of this solution, let just
 * to implement it.
 * Time complexity will be O(n*k)
 * Space complexity will be O(k) only.
NOTE END */

func Array_find_element_occur_more_nbyk_unsorted_03(arr []int, k int) string {

	possibleMajCandidates := getPossibleMajorityCandidates(arr, k)

	res := make([]string, 0)

	for _, majCandidate := range possibleMajCandidates {

		// checking the count of that element in array
		elementCnt := 0
		for _, arrElement := range arr {
			if arrElement == majCandidate.element {
				elementCnt++
			}
		}
		if elementCnt > len(arr)/k {
			res = append(res, strconv.Itoa(majCandidate.element))
		}
	}
	return strings.Join(res, ",")

}

type Pair struct {
	votes int
	element int
}

func getPossibleMajorityCandidates(arr []int, k int) []Pair {
	possibleMajorityElements := make([]Pair, k-1)

	for i := 0; i < len(arr); i++ {

		var elementFound bool // a variable to know if arr[i] is found or not in the possibleMajorityElements array

		// if arr[i] is available in possibleMajorityElements, we need to increase the voting count
		for _, majCandidate := range possibleMajorityElements {
			if majCandidate.element == arr[i] {
				majCandidate.votes++
				elementFound = true
				break
			}
		}

		if !elementFound {
			// we check if any element present in possibleMajorityElement array has count as 0,
			// we there is such element, then we insert the current element into this array.
			// Also, we decrease the votes of each element present in this array (possibleMajorityElement)

			var isNewElementFound bool  // this var keeps track if we have assigned new element a place in our possible
			// candidate of majority, in that case we do not wish to decrement the older votes count, why? coz the its the addition of new possibility of majority candidate.

			for j := 0; j < len(possibleMajorityElements); j++ {
				if possibleMajorityElements[j].votes == 0 {
					possibleMajorityElements[j].element = arr[i]
					possibleMajorityElements[j].votes = 1
					isNewElementFound = true
					break
				}
			}


			if !isNewElementFound {
				for j := 0; j < len(possibleMajorityElements); j++ {
					if possibleMajorityElements[j].votes > 0 {
						possibleMajorityElements[j].votes--
					}
				}
			}
		}
	}

	return possibleMajorityElements
}
