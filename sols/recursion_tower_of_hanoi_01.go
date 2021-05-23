package sols

import "strconv"

/* NOTE BEGIN
 * A disc can not ever sit on top of a larger disc.
 * We will try to solve this with help of recursion. So, how is this problem self similar?
 * Try to solve this when there are only 2 discs in one peg, you would be easily conclude that we need to move the
 * one disc at top to intermediate peg and then the last disc to the destination peg and finally the disc from
 * intermediate peg to destination peg.

 * Similarly, we can find the self similarity for n number of discs in the stack.
 * Move top n-1 discs from origin peg to intermediate peg.
 * Move the nth disc from origin peg to destination peg.
 * Move the n-1 discs from intermediate peg to destination peg.
 * That is all.
 * We can keep a counter whenever we are moving a disc, and finally when all discs are moved the counter will be the
 * answer.
 * One more thing, what will be the base condition here? What is the easy number of disc to move? if there is no disc
 * we know that we don't need to take any action. So, base case could be when number of disc is 0, do nothing
NOTE END */

// we assume there are 3 pegs and we will always move discs from 1st peg to 3rd peg
func Recursion_tower_of_hanoi_01(discNo int) string {

	ctr := 0
	moveDiscs(discNo, "A", "C", &ctr)
	return strconv.Itoa(ctr)
}

func moveDiscs(discNo int, initPeg, finalPeg string, ctr *int) {

	if discNo == 0 {
		return
	}
	intermediatePeg := getIntermediatePeg(initPeg, finalPeg)
	moveDiscs(discNo-1, initPeg, intermediatePeg, ctr)
	*ctr++
	moveDiscs(discNo-1, intermediatePeg, finalPeg, ctr)
}

func getIntermediatePeg(initPeg, finalPeg string) string {
	if initPeg == "A" {
		if finalPeg == "B" {
			return "C"
		} else if finalPeg == "C" {
			return "B"
		} else {
			// this should not be possible
			panic("wrong initialization of peg")
		}
	} else if initPeg == "B" {
		if finalPeg == "A" {
			return "C"
		} else {
			return "A"
		}
	} else {
		if finalPeg == "A" {
			return "B"
		} else {
			return "A"
		}
	}
}
