/**
* Problem 24: Lexicographic permutations
* @see {@link https://projecteuler.net/problem=24}
*
* A permutation is an ordered arrangement of objects. For
* example, 3124 is one possible permutation of the digits
* 1, 2, 3 and 4. If all the permutations are listed numerically
* or alphabetically, we call it lexicographic order. The
* lexicographic permutations of 0, 1 and 2 are:
* 012   021   102   120   201   210
*
* What is the millionth lexicographic permutation of the
* digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*
* @author ddaniel27
 */
package problem24

// Define factorials in advance just to avoid
// calculate them in each digit
// This MUST be a read only map
var factorials map[uint8]uint = map[uint8]uint{
	10: 3628800,
	9:  362880,
	8:  40320,
	7:  5040,
	6:  720,
	5:  120,
	4:  24,
	3:  6,
	2:  2,
	1:  1,
	0:  1,
}

func Problem24(target uint) string {
	// Edges
	if target > factorials[10] || target == 0 {
		return ""
	}

	availableDigits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	currentSum := uint(0)
	res := make([]rune, 0)

	// Iterate over each digit
	for i := uint8(1); i <= 10; i++ {
		factorial := factorials[10-i]

		// Iterate over all possible options for next digit
		for i, v := range availableDigits {
			// Add the total possibilities for current position
			currentSum += factorial

			// If we match or pass our target position
			if currentSum >= target {
				currentSum -= factorial // Go back one step
				res = append(res, v)    // Current option is the right one

				// delete it from options
				availableDigits = append(availableDigits[:i], availableDigits[i+1:]...)
				break
			}
		}
	}

	var str string
	for _, v := range res {
		str += string(v)
	}

	return str
}
